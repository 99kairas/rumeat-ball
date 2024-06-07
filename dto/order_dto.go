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
	Quantity     int       `json:"quantity"`
	PricePerItem float64   `json:"price_per_item"`
	TotalPrice   float64   `json:"total_price"`
}

type OrderResponse struct {
	ID     uuid.UUID   `json:"id"`
	Status string      `json:"status"`
	Date   string      `json:"date"`
	Total  float64     `json:"total"`
	Items  []OrderItem `json:"items"`
}

func ConvertToOrderModel(req OrderRequest) models.Order {
	return models.Order{
		ID:     uuid.New(),
		Date:   time.Now(),
		Status: "processed",
	}
}

func ConvertToOrderResponse(order models.Order, items []OrderItem) OrderResponse {
	dateFormat := time.Now().Format("02 January 2006")
	return OrderResponse{
		ID:     order.ID,
		Status: order.Status,
		Date:   dateFormat,
		Total:  order.Total,
		Items:  items,
	}
}
