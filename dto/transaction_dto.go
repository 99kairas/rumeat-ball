package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransactionRequest struct {
	OrderID     string  `json:"order_id"`
	TotalAmount float64 `json:"total_amount"`
}

type TransactionResponse struct {
	ID          uuid.UUID `json:"id"`
	OrderID     string    `json:"order_id"`
	UserID      uuid.UUID `json:"user_id"`
	Status      string    `json:"status"`
	PaymentURL  string    `json:"payment_url"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type MidTransNotification struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
}

// type TransactionRequest struct {
// 	OrderID    uuid.UUID `json:"order_id" form:"order_id"`
// 	TotalPrice float64   `json:"total_price" form:"total_price"`
// 	UserID     uuid.UUID `json:"user_id" form:"user_id"`
// }

// type TransactionResponse struct {
// 	ID         uuid.UUID `json:"id" form:"id"`
// 	OrderID    uuid.UUID `json:"order_id" form:"order_id"`
// 	TotalPrice float64   `json:"total_price" form:"total_price"`
// 	Status     string    `json:"status" form:"status"`
// 	PaymentURL string    `json:"payment_url" form:"payment_url"`
// }

// func ConvertToTransactionModel(request TransactionRequest) models.Transaction {
// 	return models.Transaction{
// 		ID:         uuid.New(),
// 		OrderID:    request.OrderID,
// 		TotalPrice: request.TotalPrice,
// 		Status:     "pending",
// 		UserID:     request.UserID,
// 	}
// }

// func ConvertToTransactionResponse(transaction models.Transaction) TransactionResponse {
// 	return TransactionResponse{
// 		ID:         transaction.ID,
// 		OrderID:    transaction.OrderID,
// 		TotalPrice: transaction.TotalPrice,
// 		Status:     transaction.Status,
// 	}
// }

// func ConvertToTransactionResponseList(request []models.Transaction) []TransactionResponse {
// 	var response []TransactionResponse
// 	for _, v := range request {
// 		response = append(response, TransactionResponse{
// 			ID:         v.ID,
// 			OrderID:    v.OrderID,
// 			TotalPrice: v.TotalPrice,
// 		})
// 	}
// 	return response
// }

// func ConvertToTransactionModelList(request []TransactionRequest) []models.Transaction {
// 	var response []models.Transaction
// 	for _, v := range request {
// 		response = append(response, ConvertToTransactionModel(v))
// 	}
// 	return response
// }
