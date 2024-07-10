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
	fmt.Print("ini adalah user id ", userID)
	// userID := "74c89fe0-0f95-40d8-9912-fd181dfaf7d5"

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
	orderData.ID = uuid.New()

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
