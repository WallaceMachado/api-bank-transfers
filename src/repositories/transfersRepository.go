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
