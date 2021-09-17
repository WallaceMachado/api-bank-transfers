package repositories

import (
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func Create(account models.Account) (uint, error) {
	repository := database.GetDatabase()
	err := repository.Create(&account).Error

	if err != nil {
		return account.ID, err
	}

	return account.ID, nil
}
