package repositories

import (
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

type TransferRepository struct {
}

func (repository TransferRepository) SaveTransfer(transfer models.Transfer) (models.Transfer, error) {
	db := database.GetDatabase()
	err := db.Create(&transfer).Error

	if err != nil {
		return transfer, err
	}

	return transfer, nil
}

func (repository TransferRepository) GetTransfersById(id string) ([]models.Transfer, error) {
	db := database.GetDatabase()

	var transfers []models.Transfer

	err := db.Where("account_origin_id = ? OR Account_destination_id = ?", id, id).Find(&transfers).Error

	if err != nil {
		return []models.Transfer{}, err
	}

	return transfers, nil
}
