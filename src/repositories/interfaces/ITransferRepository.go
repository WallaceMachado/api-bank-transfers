package interfaces

import "github.com/wallacemachado/api-bank-transfers/src/models"

type ITransferRepository interface {
	SaveTransfer(transfer *models.Transfer) (*models.Transfer, error)
	GetTransfersById(id string) ([]models.Transfer, error)
}
