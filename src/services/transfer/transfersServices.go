package services

import (
	"errors"

	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories/interfaces"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
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

	account_origin, _ := s.repositoryAccount.GetAccountById(transfer.Account_origin_id)

	account_destination, _ := s.repositoryAccount.GetAccountById(transfer.Account_destination_id)

	if len(account_destination.ID) == 0 || len(account_origin.ID) == 0 {
		return &models.Transfer{}, errors.New("Account not found")
	}

	if account_origin.Balance < transfer.Amount {
		return &models.Transfer{}, errors.New("Insufficient balance")
	}

	newTransfer, err := s.repositoryTransfer.SaveTransfer(transfer)
	if err != nil {
		return &models.Transfer{}, err
	}

	account_origin.Balance = account_origin.Balance - transfer.Amount

	_, err = s.repositoryAccount.UpdateBalanceAccount(account_origin)
	if err != nil {
		return &models.Transfer{}, err
	}

	account_destination.Balance = account_destination.Balance + transfer.Amount

	_, err = s.repositoryAccount.UpdateBalanceAccount(account_destination)
	if err != nil {
		return &models.Transfer{}, err
	}

	return newTransfer, nil

}

func (s *TransferService) ListAllTransfersByAccount(id string) (responses.ResponseTransfersByAccount, error) {

	transfers, err := s.repositoryTransfer.GetTransfersById(id)
	if err != nil {
		return responses.ResponseTransfersByAccount{}, err
	}

	var transferResponse responses.ResponseTransfersByAccount

	for _, t := range transfers {
		if t.Account_origin_id == id {
			transferResponse.TranfersSent = append(transferResponse.TranfersSent, t)
		} else {

			transferResponse.TranfersReceived = append(transferResponse.TranfersReceived, t)
		}
	}

	return transferResponse, nil

}
