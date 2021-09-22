package services

import (
	"errors"

	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories/interfaces"
)

type AccountService struct {
	repository interfaces.IAccountRepository
}

func NewAccountService(repo interfaces.IAccountRepository) *AccountService {
	return &AccountService{
		repository: repo,
	}
}

func (s *AccountService) CreateAccount(account models.Account) (string, error) {

	account, err := s.repository.GetAccountByCpf(account.Cpf)
	if err != nil {
		return "", errors.New("Error in cpf validation")
	}

	if account.Cpf != "" {
		return "", errors.New("CPF already exists")
	}

	return s.repository.CreateAccount(account)
}

func (s *AccountService) ListAllAccounts() ([]models.Account, error) {

	return s.repository.GetAllAccounts()
}

func (s *AccountService) GetBalance(id string) (float64, error) {

	account, err := s.repository.GetAccountById(id)
	if err != nil {
		return 0, err
	}

	balance := account.Balance

	return balance, nil
}
