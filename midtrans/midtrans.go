package midtrans

// func CreatePayment(transaction models.Transaction) (string, error) {
// 	midtransClient := GetMidtransClient()

// 	chargeReq := &midtrans.ChargeReq{
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  transaction.OrderID.String(),
// 			GrossAmt: int64(transaction.TotalPrice),
// 		},
// 		CustomerDetails: &midtrans.CustomerDetails{
// 			FName: "FirstName",
// 			LName: "LastName",
// 			Email: "customer@example.com",
// 			Phone: "081234567890",
// 		},
// 		CreditCard: &midtrans.CreditCardDetails{
// 			Secure: true,
// 		},
// 	}

// 	chargeResponse, err := midtransClient.ChargeTransaction(chargeReq)
// 	if err != nil {
// 		return "", err
// 	}

// 	return chargeResponse.RedirectURL, nil
// }
