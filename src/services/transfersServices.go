package services

import (
	"errors"

	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
)

func CreateTransfer(transfer models.Transfer) (models.Transfer, error) {

	account_origin, err := repositories.GetAccountById(transfer.Account_origin_id)
	if err != nil {
		return models.Transfer{}, err
	}

	account_destination, err := repositories.GetAccountById(transfer.Account_destination_id)
	if err != nil {
		return models.Transfer{}, err
	}

	if len(account_destination.ID) == 0 || len(account_origin.ID) == 0 {
		return models.Transfer{}, errors.New("account not found")
	}

	if account_origin.Balance < transfer.Amount {
		return models.Transfer{}, errors.New("insufficient balance")
	}

	newTransfer, err := repositories.SaveTransfer(transfer)
	if err != nil {
		return models.Transfer{}, err
	}

	account_origin.Balance = account_origin.Balance - transfer.Amount

	_, err = repositories.UpdateBalanceAccount(account_origin)
	if err != nil {
		return models.Transfer{}, err
	}

	account_destination.Balance = account_destination.Balance + transfer.Amount

	_, err = repositories.UpdateBalanceAccount(account_destination)
	if err != nil {
		return models.Transfer{}, err
	}

	return newTransfer, nil

}
