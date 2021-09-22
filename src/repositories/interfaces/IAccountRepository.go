package interfaces

import "github.com/wallacemachado/api-bank-transfers/src/models"

type IAccountRepository interface {
	CreateAccount(account models.Account) (string, error)
	GetAllAccounts() ([]models.Account, error)
	GetAccountById(id string) (models.Account, error)
	GetAccountByCpf(cpf string) (models.Account, error)
	UpdateBalanceAccount(account models.Account) (models.Account, error)
}
