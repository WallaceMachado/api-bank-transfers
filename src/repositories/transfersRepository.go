package repositories

import (
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func SaveTransfer(transfer models.Transfer) (models.Transfer, error) {
	repository := database.GetDatabase()
	err := repository.Create(&transfer).Error

	if err != nil {
		return transfer, err
	}

	return transfer, nil
}

func GetTransfersById(id string) ([]models.Transfer, error) {
	repository := database.GetDatabase()

	var transfers []models.Transfer

	err := repository.Where("account_origin_id = ? OR Account_destination_id = ?", id, id).Find(&transfers).Error

	if err != nil {
		return []models.Transfer{}, err
	}

	return transfers, nil
}
