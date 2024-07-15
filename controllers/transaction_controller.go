package controllers

import (
	"net/http"
	"rumeat-ball/dto"
	m "rumeat-ball/middlewares"
	"rumeat-ball/models"
	"rumeat-ball/repositories"
	"rumeat-ball/util"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateTransactionController(c echo.Context) error {
	userID := m.ExtractTokenUserId(c)
	if userID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	var req dto.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: err.Error(),
		})
	}

	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error fetching user",
			Response: err.Error(),
		})
	}

	transaction := models.Transaction{
		ID:         uuid.New(),
		OrderID:    req.OrderID,
		UserID:     userID,
		Status:     "pending",
		TotalPrice: req.TotalAmount,
	}

	snapResp, err := util.GetPaymentURL(&transaction, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error creating midtrans transaction",
			Response: err.Error(),
		})
	}

	transaction.PaymentURL = snapResp.RedirectURL
	transaction, err = repositories.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error creating transaction",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "success create transaction",
		Response: dto.TransactionResponse{
			ID:          transaction.ID,
			OrderID:     transaction.OrderID,
			UserID:      transaction.UserID,
			Status:      transaction.Status,
			PaymentURL:  transaction.PaymentURL,
			TotalAmount: transaction.TotalPrice,
			CreatedAt:   transaction.CreatedAt,
		},
	})
}

func HandleMidTransNotificationController(c echo.Context) error {
	var notification dto.MidTransNotification
	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: err.Error(),
		})
	}

	_, err := repositories.GetTransactionByOrderID(notification.OrderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.Response{
			Message:  "transaction not found",
			Response: err.Error(),
		})
	}

	if _, err := repositories.UpdateTransactionStatus(notification.OrderID, notification.TransactionStatus); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error updating transaction status",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success handle midtrans notification",
	})
}

// func CreateTransactionController(c echo.Context) error {
// 	var transactionReq dto.TransactionRequest
// 	if err := c.Bind(&transactionReq); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.Response{
// 			Message:  "Invalid request",
// 			Response: err.Error(),
// 		})
// 	}

// 	transaction := dto.ConvertToTransactionModel(transactionReq)
// 	createdTransaction, err := repositories.CreateTransaction(transaction)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.Response{
// 			Message:  "Failed to create transaction",
// 			Response: err.Error(),
// 		})
// 	}

// 	paymentURL, err := midtrans.CreatePayment(createdTransaction)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.Response{
// 			Message:  "Failed to create payment",
// 			Response: err.Error(),
// 		})
// 	}

// 	transactionResponse := dto.ConvertToTransactionResponse(createdTransaction)
// 	transactionResponse.PaymentURL = paymentURL

// 	return c.JSON(http.StatusOK, dto.Response{
// 		Message:  "Transaction created",
// 		Response: transactionResponse,
// 	})
// }

// func PaymentNotificationController(c echo.Context) error {
// 	var notification midtrans.TransactionStatusResponse
// 	if err := c.Bind(&notification); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.Response{
// 			Message:  "Invalid request",
// 			Response: err.Error(),
// 		})
// 	}

// 	transactionIDStr := notification.OrderID
// 	status := notification.TransactionStatus

// 	transactionID, err := uuid.Parse(transactionIDStr)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.Response{
// 			Message:  "Invalid transaction ID",
// 			Response: err.Error(),
// 		})
// 	}

// 	if err := repositories.UpdateTransactionStatus(transactionID, status); err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.Response{
// 			Message:  "Failed to update transaction status",
// 			Response: err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, dto.Response{
// 		Message:  "Transaction status updated",
// 		Response: nil,
// 	})
// }
