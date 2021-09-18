package services

import (
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
)

func CreateAccount(account models.Account) (uint, error) {

	return repositories.Create(account)
}

func ListAllAccounts() ([]responses.ResponseGetAccount, error) {

	return repositories.GetAll()
}

func GetBalance(id int) (float32, error) {

	account, err := repositories.GetAccountById(id)
	if err != nil {
		return 0, err
	}

	balance := account.Balance

	return balance, nil
}
