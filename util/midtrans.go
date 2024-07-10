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

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReg := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.OrderID.String(),
			GrossAmt: int64(transaction.TotalPrice),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReg)
	if err != nil {
		return snapTokenResp, err
	}

	return snapTokenResp, nil
}
