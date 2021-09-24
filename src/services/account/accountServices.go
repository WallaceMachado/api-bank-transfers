package services

import (
	"errors"
	"fmt"

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

func (s *AccountService) CreateAccount(acc *models.Account) (string, error) {

	account, err := s.repository.GetAccountByCpf(acc.Cpf)
	fmt.Println("cpf: ", err)
	if err != nil {
		return "", err
	}

	if account.Cpf != "" {
		return "", errors.New("CPF already exists")
	}

	return s.repository.CreateAccount(acc)
}

func (s *AccountService) ListAllAccounts() ([]models.Account, error) {

	return s.repository.GetAllAccounts()
}

func (s *AccountService) GetBalance(id string) (float64, error) {

	account, err := s.repository.GetAccountById(id)
	if err != nil {
		return 0, err
	}

	if account.Cpf == "" {
		return 0, errors.New("Non-existent account")
	}

	balance := account.Balance

	return balance, nil
}
