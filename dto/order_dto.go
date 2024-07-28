package dto

import (
	"rumeat-ball/models"
	"time"

	"github.com/google/uuid"
)

type OrderRequest struct {
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	MenuID       uuid.UUID `json:"menu_id"`
	UserID       uuid.UUID `json:"user_id"`
	Quantity     int       `json:"quantity"`
	PricePerItem float64   `json:"price_per_item"`
	TotalPrice   float64   `json:"total_price"`
}

type OrderResponse struct {
	ID     string      `json:"id"`
	UserID uuid.UUID   `json:"user_id"`
	Status string      `json:"status"`
	Date   string      `json:"date"`
	Total  float64     `json:"total"`
	Items  []OrderItem `json:"items"`
}

type AdminOrderResponse struct {
	ID       string              `json:"id"`
	UserID   uuid.UUID           `json:"user_id"`
	UserName string              `json:"user_name"`
	Status   string              `json:"status"`
	Date     string              `json:"date"`
	Total    float64             `json:"total"`
	Items    []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	MenuID       uuid.UUID `json:"menu_id"`
	Quantity     int       `json:"quantity"`
	PricePerItem float64   `json:"price_per_item"`
	TotalPrice   float64   `json:"total_price"`
}

func ConvertToOrderModel(req OrderRequest, userID uuid.UUID) models.Order {
	return models.Order{
		UserID: userID,
		Date:   time.Now(),
		Status: "cart",
	}
}

func ConvertToOrderResponse(order models.Order, items []OrderItem, userID uuid.UUID) OrderResponse {
	dateFormat := order.Date.Format("02 January 2006 15:04")
	return OrderResponse{
		ID:     order.ID,
		UserID: userID,
		Status: order.Status,
		Date:   dateFormat,
		Total:  order.Total,
		Items:  items,
	}
}

func ConvertToOrderRequest(order models.Order) OrderRequest {
	return OrderRequest{
		Items: []OrderItem{},
	}
}

func ConvertToAdminOrderResponse(order models.Order, items []models.DetailOrder, userName string) AdminOrderResponse {
	var itemResponses []OrderItemResponse
	for _, item := range items {
		itemResponses = append(itemResponses, OrderItemResponse{
			MenuID:       item.MenuID,
			Quantity:     item.Quantity,
			PricePerItem: item.TotalPrice / float64(item.Quantity),
			TotalPrice:   item.TotalPrice,
		})
	}

	return AdminOrderResponse{
		ID:       order.ID,
		UserID:   order.UserID,
		UserName: userName,
		Status:   order.Status,
		Date:     order.Date.Format("02 January 2006 15:04"),
		Total:    order.Total,
		Items:    itemResponses,
	}
}
