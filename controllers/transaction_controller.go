package controllers

// "rumeat-ball/midtrans"

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
