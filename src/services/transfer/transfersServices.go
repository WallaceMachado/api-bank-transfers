package services

import (
	"errors"

	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories/interfaces"
)

type TransferService struct {
	repositoryAccount  interfaces.IAccountRepository
	repositoryTransfer interfaces.ITransferRepository
}

func NewTransferService(repoAcc interfaces.IAccountRepository, repoTransfer interfaces.ITransferRepository) *TransferService {
	return &TransferService{
		repositoryAccount:  repoAcc,
		repositoryTransfer: repoTransfer,
	}
}

func (s *TransferService) CreateTransfer(transfer *models.Transfer) (*models.Transfer, error) {

	account_origin, err := s.repositoryAccount.GetAccountById(transfer.Account_origin_id)
	if err != nil {
		return &models.Transfer{}, err
	}

	if len(account_origin.ID) == 0 {
		return &models.Transfer{}, errors.New("Origin account not found")
	}

	account_destination, err := s.repositoryAccount.GetAccountById(transfer.Account_destination_id)
	if err != nil {
		return &models.Transfer{}, err
	}

	if len(account_destination.ID) == 0 {
		return &models.Transfer{}, errors.New("Destination account not found")
	}

	if account_origin.Balance < transfer.Amount {
		return &models.Transfer{}, errors.New("Insufficient balance")
	}

	newTransfer, err := s.repositoryTransfer.CreateTransfer(transfer)
	if err != nil {
		return &models.Transfer{}, err
	}

	account_origin.Balance = account_origin.Balance - transfer.Amount

	_, err = s.repositoryAccount.UpdateAccount(account_origin)
	if err != nil {
		return &models.Transfer{}, err
	}

	account_destination.Balance = account_destination.Balance + transfer.Amount

	_, err = s.repositoryAccount.UpdateAccount(account_destination)
	if err != nil {
		return &models.Transfer{}, err
	}

	return newTransfer, nil

}

func (s *TransferService) ListAllTransfersByAccount(id string) ([]models.Transfer, error) {

	account, err := s.repositoryAccount.GetAccountById(id)
	if err != nil {
		return []models.Transfer{}, err
	}

	if len(account.ID) == 0 {
		return []models.Transfer{}, errors.New("Account not found")
	}

	transfers, err := s.repositoryTransfer.GetTransfersByAccountId(id)
	if err != nil {
		return []models.Transfer{}, err
	}

	return transfers, nil

}
