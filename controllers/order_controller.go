package controllers

import (
	"fmt"
	"net/http"
	"rumeat-ball/dto"
	m "rumeat-ball/middlewares"
	"rumeat-ball/models"
	"rumeat-ball/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateOrderController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)

	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	var orderReq = dto.OrderRequest{}
	errBind := c.Bind(&orderReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: errBind.Error(),
		})
	}
	orderData := dto.ConvertToOrderModel(orderReq, userID)

	// Calculate total price
	var totalOrderPrice float64
	var orderItems []dto.OrderItem

	for _, item := range orderReq.Items {
		menu, err := repositories.GetMenuByID(item.MenuID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.Response{
				Message:  "error fetching menu data",
				Response: err.Error(),
			})
		}

		totalItemPrice := float64(item.Quantity) * menu.Price
		totalOrderPrice += totalItemPrice
		fmt.Print(userID)

		// Add item to order items for response
		orderItems = append(orderItems, dto.OrderItem{
			MenuID:       item.MenuID,
			UserID:       userID,
			Quantity:     item.Quantity,
			PricePerItem: menu.Price,
			TotalPrice:   totalItemPrice,
		})
	}

	orderData.Total = totalOrderPrice

	// Create order
	order, err := repositories.CreateOrder(orderData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error create order",
			Response: err.Error(),
		})
	}

	// Create detail orders
	for _, item := range orderItems { // Use orderItems to ensure PricePerItem is set
		detailOrder := models.DetailOrder{
			ID:         uuid.New(),
			OrderID:    order.ID,
			MenuID:     item.MenuID,
			Quantity:   item.Quantity,
			TotalPrice: item.TotalPrice, // Calculate total price based on the Menu price and quantity
		}
		_, err = repositories.CreateDetailOrder(detailOrder)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.Response{
				Message:  "error create detail order",
				Response: err.Error(),
			})
		}
	}

	// Convert To Response
	orderResponse := dto.ConvertToOrderResponse(order, orderItems, userID)

	return c.JSON(http.StatusCreated, dto.Response{
		Message:  "success create order",
		Response: orderResponse,
	})
}

func GetAllOrdersController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	orders, err := repositories.GetOrdersByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error fetching order data",
			Response: err.Error(),
		})
	}

	// Convert To Response
	var orderResponses []dto.OrderResponse
	for _, order := range orders {
		orderItems, err := repositories.GetOrderItemsByOrderID(order.ID, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.Response{
				Message:  "error fetching order items data",
				Response: err.Error(),
			})
		}
		orderResponse := dto.ConvertToOrderResponse(order, orderItems, userID)
		orderResponses = append(orderResponses, orderResponse)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success get all orders",
		Response: map[string]any{
			"orders": orderResponses,
		},
	})
}

func GetOrderByIDController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	orderID := c.Param("id")
	if orderID == "" {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error parse id",
			Response: "order id is invalid",
		})
	}

	order, err := repositories.GetOrderByOrderID(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error fetching order data",
			Response: err.Error(),
		})
	}

	// Convert To Response
	orderItems, err := repositories.GetOrderItemsByOrderID(order.ID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error fetching order items data",
			Response: err.Error(),
		})
	}
	orderResponse := dto.ConvertToOrderResponse(order, orderItems, userID)

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get order by id",
		Response: orderResponse,
	})
}

func CancelOrderController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error parse id",
			Response: err.Error(),
		})
	}

	err = repositories.CancelOrder(orderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error cancel order",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success cancel order",
		Response: nil,
	})
}
