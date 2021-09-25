package responses

import "github.com/wallacemachado/api-bank-transfers/src/models"

type ResponseTransfersByAccount struct {
	TranfersSent     []models.Transfer `json:"tranfersSent"`
	TranfersReceived []models.Transfer `json:"tranfersReceived "`
}
