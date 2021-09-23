package services_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/database/migrations"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	servicesAccount "github.com/wallacemachado/api-bank-transfers/src/services/account"
	servicesTransfer "github.com/wallacemachado/api-bank-transfers/src/services/transfer"
	"github.com/wallacemachado/api-bank-transfers/src/utils/dtos"
)

func TestLogin(t *testing.T) {

	database.StartDatabase("test")

	db := database.GetDatabase()

	defer database.CloseConn()

	repositoryAccount := &repositories.AccountRepository{}

	repositoryTransfer := &repositories.TransferRepository{}

	serviceAccount := servicesAccount.NewAccountService(repositoryAccount)

	servicesTransfer := servicesTransfer.NewTransferService(repositoryAccount, repositoryTransfer)

	//cpf gerado aleatoriamente no site: https://www.4devs.com.br/gerador_de_cpf
	createAccountOrigin := &dtos.CreateAccountDTO{
		Cpf:     "16255912094",
		Name:    "account1",
		Secret:  "123456",
		Balance: 1000,
	}

	createAccountDestination := &dtos.CreateAccountDTO{
		Cpf:     "05439535055",
		Name:    "account2",
		Secret:  "123456",
		Balance: 1000,
	}

	accountOrigin, _ := models.NewAccount(
		createAccountOrigin.Name,
		createAccountOrigin.Cpf,
		createAccountOrigin.Secret,
		createAccountOrigin.Balance,
	)
	accountDestination, _ := models.NewAccount(
		createAccountDestination.Name,
		createAccountDestination.Cpf,
		createAccountDestination.Secret,
		createAccountDestination.Balance,
	)

	accountOriginId, err := serviceAccount.CreateAccount(accountOrigin)
	accountOrigin.ID = accountOriginId

	accountDestinationId, _ := serviceAccount.CreateAccount(accountDestination)
	accountDestination.ID = accountDestinationId

	transfer, _ := models.NewTransfer(accountOriginId, accountDestinationId, 500)

	result, err := servicesTransfer.CreateTransfer(transfer)
	require.Nil(t, err)
	assert.NotEmpty(t, result.ID)

	accountInvalidID := uuid.NewV4().String()

	transferInvalidOrigin, _ := models.NewTransfer(accountInvalidID, accountDestinationId, 500)
	result, err = servicesTransfer.CreateTransfer(transferInvalidOrigin)
	require.Error(t, err)
	require.EqualError(t, err, "Account not found")
	assert.Empty(t, result.ID)

	transferInvalidDestination, _ := models.NewTransfer(accountOriginId, accountInvalidID, 500)
	result, err = servicesTransfer.CreateTransfer(transferInvalidDestination)
	require.Error(t, err)
	require.EqualError(t, err, "Account not found")
	assert.Empty(t, result.ID)

	transferInvalidAmount, _ := models.NewTransfer(accountOriginId, accountDestinationId, 1001)
	result, err = servicesTransfer.CreateTransfer(transferInvalidAmount)
	require.Error(t, err)
	require.EqualError(t, err, "Insufficient balance")
	assert.Empty(t, result.ID)

	migrations.DeleteTablesTestDb(db)

}
