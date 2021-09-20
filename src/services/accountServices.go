package services

import (
	"errors"

	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
)

func CreateAccount(account models.Account) (string, error) {

	account, err := repositories.GetAccountByCpf(account.Cpf)
	if err != nil {
		return "", errors.New("Error in cpf validation")
	}

	if account.Cpf != "" {
		return "", errors.New("CPF already exists")
	}

	return repositories.CreateAccount(account)
}

func ListAllAccounts() ([]models.Account, error) {

	return repositories.GetAllAccounts()
}

func GetBalance(id string) (float64, error) {

	account, err := repositories.GetAccountById(id)
	if err != nil {
		return 0, err
	}

	balance := account.Balance

	return balance, nil
}
