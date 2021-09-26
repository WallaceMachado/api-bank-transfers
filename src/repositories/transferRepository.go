package repositories

import (
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/shared/database"
)

type TransferRepository struct {
}

func (repository TransferRepository) CreateTransfer(transfer *models.Transfer) (*models.Transfer, error) {
	db := database.GetDatabase()
	err := db.Create(&transfer).Error

	if err != nil {
		return &models.Transfer{}, err
	}

	return transfer, nil
}

func (repository TransferRepository) GetTransfersByAccountId(id string) ([]models.Transfer, error) {
	db := database.GetDatabase()

	var transfers []models.Transfer

	err := db.Where("account_origin_id = ? OR Account_destination_id = ?", id, id).Find(&transfers).Error

	if err != nil {
		return []models.Transfer{}, err
	}

	return transfers, nil
}
