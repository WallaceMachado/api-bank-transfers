package services

import (
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
)

func CreateAccount(account models.Account) (string, error) {

	return repositories.Create(account)
}

func ListAllAccounts() ([]models.Account, error) {

	return repositories.GetAll()
}

func GetBalance(id int) (float64, error) {

	account, err := repositories.GetAccountById(id)
	if err != nil {
		return 0, err
	}

	balance := account.Balance

	return balance, nil
}
