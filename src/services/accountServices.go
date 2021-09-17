package services

import (
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
)

func CreateAccount(account models.Account) (uint, error) {

	return repositories.Create(account)
}
