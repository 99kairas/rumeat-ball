package util

import (
	"os"
	"rumeat-ball/models"

	"github.com/veritrans/go-midtrans"
)

func GetPaymentURL(transaction *models.Transaction, user *models.User) (midtrans.SnapResponse, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("Server_Key")
	midclient.ClientKey = os.Getenv("Client_Key")
	midclient.APIEnvType = midtrans.Sandbox

	// fmt.Println("Server Key : ", configs.MIDTRANS_SERVER_KEY)
	// fmt.Println("Client Key : ", configs.MIDTRANS_CLIENT_KEY)

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderID,
			GrossAmt: int64(transaction.TotalPrice),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return midtrans.SnapResponse{}, err
	}

	return snapTokenResp, nil
}
