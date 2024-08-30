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

	// Bind data notifikasi dari Midtrans
	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "error bind data",
			Response: err.Error(),
		})
	}

	// Dapatkan transaksi berdasarkan OrderID dari notifikasi
	transaction, err := repositories.GetTransactionByOrderID(notification.OrderID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.Response{
			Message:  "transaction not found",
			Response: err.Error(),
		})
	}

	// Periksa status transaksi yang diterima dari Midtrans
	var newStatus string
	switch notification.TransactionStatus {
	case "settlement", "capture":
		newStatus = "successed"
	case "pending":
		newStatus = "pending"
	case "deny", "cancel", "expire":
		newStatus = "failed"
	default:
		return c.JSON(http.StatusBadRequest, dto.Response{
			Message:  "invalid transaction status",
			Response: "unexpected transaction status from Midtrans",
		})
	}

	// Memperbarui status transaksi di database
	updatedTransaction, err := repositories.UpdateTransactionStatus(transaction.OrderID, newStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error updating transaction status",
			Response: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "success handle midtrans notification",
		Response: dto.TransactionResponse{
			ID:          updatedTransaction.ID,
			OrderID:     updatedTransaction.OrderID,
			UserID:      updatedTransaction.UserID,
			Status:      updatedTransaction.Status,
			PaymentURL:  updatedTransaction.PaymentURL,
			TotalAmount: updatedTransaction.TotalPrice,
			CreatedAt:   updatedTransaction.CreatedAt,
		},
	})
}

func AdminGetAllTransactionsController(c echo.Context) error {
	adminID := m.ExtractTokenUserId(c)

	if adminID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, dto.Response{
			Message:  "unauthorized",
			Response: "permission denied: user is not valid",
		})
	}

	userName := c.QueryParam("name")
	userIDStr := c.QueryParam("user_id")
	transactionIDStr := c.QueryParam("transaction_id")
	orderID := c.QueryParam("order_id")

	var transactions []models.Transaction
	var err error

	switch {
	case userIDStr != "":
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.Response{
				Message:  "invalid user_id format",
				Response: err.Error(),
			})
		}
		transactions, err = repositories.GetTransactionsByUserID(userID)
	case transactionIDStr != "":
		transactionID, err := uuid.Parse(transactionIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.Response{
				Message:  "invalid transaction_id format",
				Response: err.Error(),
			})
		}
		transaction, err := repositories.GetTransactionByID(transactionID)
		if err == nil {
			transactions = append(transactions, transaction)
		}
	case userName != "":
		transactions, err = repositories.GetTransactionsByUserName(userName)
	case orderID != "":
		transactions, err = repositories.GetTransactionsByOrderID(orderID)
	default:
		transactions, err = repositories.GetAllTransactions()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Response{
			Message:  "error fetching transaction data",
			Response: err.Error(),
		})
	}

	var transactionResponses []dto.AdminTransactionResponse
	for _, transaction := range transactions {
		transactionResponse := dto.ConvertToAdminTransactionResponse(transaction, transaction.User.Name)
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message:  "success get all transactions",
		Response: transactionResponses,
	})
}
