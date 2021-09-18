package repositories

import (
	"fmt"

	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
)

func Create(account models.Account) (uint, error) {
	repository := database.GetDatabase()
	err := repository.Create(&account).Error

	if err != nil {
		return account.ID, err
	}

	return account.ID, nil
}

func GetAll() ([]responses.ResponseGetAccount, error) {
	repository := database.GetDatabase()

	var accounts []responses.ResponseGetAccount

	err := repository.Table("accounts").Select("id", "name", "cpf", "balance", "created_at").Scan(&accounts).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(accounts)

	return accounts, nil
}
