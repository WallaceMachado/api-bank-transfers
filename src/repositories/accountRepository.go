package repositories

import (
	"fmt"

	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

type AccountRepository struct {
}

func (repository AccountRepository) CreateAccount(account *models.Account) (string, error) {
	db := database.GetDatabase()
	err := db.Create(&account).Error

	if err != nil {
		return "", err
	}

	return account.ID, nil
}

func (repository AccountRepository) GetAllAccounts() ([]models.Account, error) {
	db := database.GetDatabase()

	var accounts []models.Account

	err := db.Table("accounts").Select("id", "name", "cpf", "balance", "created_at").Scan(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repository AccountRepository) GetAccountById(id string) (models.Account, error) {
	db := database.GetDatabase()

	var account models.Account

	err := db.Where("id =?", id).First(&account).Error

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func (repository AccountRepository) GetAccountByCpf(cpf string) (models.Account, error) {
	db := database.GetDatabase()

	var account models.Account

	if err := db.Where("cpf =?", cpf).First(&account).Error; err != nil && err.Error() != "record not found" {
		fmt.Println(err.Error())
		return models.Account{}, err
	}

	return account, nil
}

func (repository AccountRepository) UpdateBalanceAccount(account models.Account) (models.Account, error) {
	db := database.GetDatabase()

	err := db.Save(&account).Error

	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}
