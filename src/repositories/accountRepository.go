package repositories

import (
	"fmt"

	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func Create(account models.Account) (string, error) {
	repository := database.GetDatabase()
	err := repository.Create(&account).Error

	if err != nil {
		return account.ID, err
	}

	return account.ID, nil
}

func GetAll() ([]models.Account, error) {
	repository := database.GetDatabase()

	var accounts []models.Account

	err := repository.Table("accounts").Select("id", "name", "cpf", "balance", "created_at").Scan(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountById(id int) (models.Account, error) {
	repository := database.GetDatabase()

	var account models.Account

	err := repository.First(&account, id).Error

	if err != nil {
		return models.Account{}, err
	}

	fmt.Println(account)

	return account, nil
}
