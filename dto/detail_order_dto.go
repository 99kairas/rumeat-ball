package dto

import (
	"rumeat-ball/models"

	"github.com/google/uuid"
)

type DetailOrderRequest struct {
	OrderID    string    `json:"order_id" form:"order_id"`
	MenuID     uuid.UUID `json:"menu_id" form:"menu_id"`
	Quantity   int       `json:"quantity" form:"quantity"`
	TotalPrice float64   `json:"total_price" form:"total_price"`
}

type DetailOrderResponse struct {
	ID         uuid.UUID `json:"id" form:"id"`
	OrderID    string    `json:"order_id" form:"order_id"`
	MenuID     uuid.UUID `json:"menu_id" form:"menu_id"`
	Quantity   int       `json:"quantity" form:"quantity"`
	TotalPrice float64   `json:"total_price" form:"total_price"`
}

func ConvertToDetailOrderModel(request DetailOrderRequest) models.DetailOrder {
	return models.DetailOrder{
		ID:         uuid.New(),
		OrderID:    request.OrderID,
		MenuID:     request.MenuID,
		Quantity:   request.Quantity,
		TotalPrice: request.TotalPrice,
	}
}

func ConvertToDetailOrderResponse(request DetailOrderResponse) DetailOrderResponse {
	return DetailOrderResponse{
		ID:         request.ID,
		OrderID:    request.OrderID,
		MenuID:     request.MenuID,
		Quantity:   request.Quantity,
		TotalPrice: request.TotalPrice,
	}
}
