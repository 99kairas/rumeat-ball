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
	ID     uuid.UUID   `json:"id"`
	UserID uuid.UUID   `json:"user_id"`
	Status string      `json:"status"`
	Date   string      `json:"date"`
	Total  float64     `json:"total"`
	Items  []OrderItem `json:"items"`
}

func ConvertToOrderModel(req OrderRequest, userID uuid.UUID) models.Order {
	return models.Order{
		ID:     uuid.New(),
		UserID: userID,
		Date:   time.Now(),
		Status: "cart",
	}
}

func ConvertToOrderResponse(order models.Order, items []OrderItem, userID uuid.UUID) OrderResponse {
	dateFormat := order.Date.Format("02 January 2006 15:04") // Using order.Date instead of time.Now()
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
