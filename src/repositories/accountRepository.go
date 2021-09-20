package repositories

import (
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func CreateAccount(account models.Account) (string, error) {
	repository := database.GetDatabase()
	err := repository.Create(&account).Error

	if err != nil {
		return "", err
	}

	return account.ID, nil
}

func GetAllAccounts() ([]models.Account, error) {
	repository := database.GetDatabase()

	var accounts []models.Account

	err := repository.Table("accounts").Select("id", "name", "cpf", "balance", "created_at").Scan(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountById(id string) (models.Account, error) {
	repository := database.GetDatabase()

	var account models.Account

	err := repository.Where("id =?", id).First(&account).Error

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func GetAccountByCpf(cpf string) (models.Account, error) {
	repository := database.GetDatabase()

	var account models.Account

	err := repository.Where("cpf =?", cpf).First(&account).Error

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func UpdateBalanceAccount(account models.Account) (models.Account, error) {
	repository := database.GetDatabase()

	err := repository.Save(&account).Error

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}
