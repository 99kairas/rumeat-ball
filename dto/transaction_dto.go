package dto

import (
	"rumeat-ball/models"
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

type AdminTransactionResponse struct {
	ID         uuid.UUID `json:"id"`
	OrderID    string    `json:"order_id"`
	TotalPrice float64   `json:"total_price"`
	PaymentURL string    `json:"payment_url"`
	Status     string    `json:"status"`
	UserID     uuid.UUID `json:"user_id"`
	UserName   string    `json:"user_name"`
	Date       string    `json:"date"`
}

func ConvertToAdminTransactionResponse(transaction models.Transaction, userName string) AdminTransactionResponse {
	return AdminTransactionResponse{
		ID:         transaction.ID,
		OrderID:    transaction.OrderID,
		TotalPrice: transaction.TotalPrice,
		PaymentURL: transaction.PaymentURL,
		Status:     transaction.Status,
		UserID:     transaction.UserID,
		UserName:   userName,
		Date:       transaction.CreatedAt.Format("02 January 2006 15:04"),
	}
}
