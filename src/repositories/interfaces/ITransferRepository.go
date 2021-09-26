package interfaces

import "github.com/wallacemachado/api-bank-transfers/src/models"

type ITransferRepository interface {
	CreateTransfer(transfer *models.Transfer) (*models.Transfer, error)
	GetTransfersByAccountId(id string) ([]models.Transfer, error)
}
