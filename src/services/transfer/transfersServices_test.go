package services_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/config"
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	servicesAccount "github.com/wallacemachado/api-bank-transfers/src/services/account"
	servicesTransfer "github.com/wallacemachado/api-bank-transfers/src/services/transfer"
	"github.com/wallacemachado/api-bank-transfers/src/utils/dtos"
)

func createTestDB() {
	config.Init()
	db := database.StartDatabase(config.DBName)
	db.Exec("DROP DATABASE IF EXISTS test;")
	db.Exec("CREATE DATABASE test;")

}

func TestCreateTransfer(t *testing.T) {
	createTestDB()
	database.StartDatabase("test")
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

	accountOriginId, _ := serviceAccount.CreateAccount(accountOrigin)
	accountOrigin.ID = accountOriginId

	accountDestinationId, _ := serviceAccount.CreateAccount(accountDestination)
	accountDestination.ID = accountDestinationId

	accountInvalidID := uuid.NewV4().String()

	t.Run("success", func(t *testing.T) {
		transfer, _ := models.NewTransfer(accountOriginId, accountDestinationId, 500)

		result, err := servicesTransfer.CreateTransfer(transfer)
		require.Nil(t, err)
		assert.NotEmpty(t, result.ID)
	})

	t.Run("Error: Origin account not found", func(t *testing.T) {

		transferInvalidOrigin, _ := models.NewTransfer(accountInvalidID, accountDestinationId, 500)
		result, err := servicesTransfer.CreateTransfer(transferInvalidOrigin)
		require.Error(t, err)
		require.EqualError(t, err, "Origin account not found")
		assert.Empty(t, result.ID)
	})

	t.Run("Error: Destination account not found", func(t *testing.T) {
		transferInvalidDestination, _ := models.NewTransfer(accountOriginId, accountInvalidID, 500)
		result, err := servicesTransfer.CreateTransfer(transferInvalidDestination)
		require.Error(t, err)
		require.EqualError(t, err, "Destination account not found")
		assert.Empty(t, result.ID)
	})

	t.Run("Error: Insufficient balance", func(t *testing.T) {
		transferInvalidAmount, _ := models.NewTransfer(accountOriginId, accountDestinationId, 1001)
		result, err := servicesTransfer.CreateTransfer(transferInvalidAmount)
		require.Error(t, err)
		require.EqualError(t, err, "Insufficient balance")
		assert.Empty(t, result.ID)
	})

}

func TestListAllTransfersByAccount(t *testing.T) {
	createTestDB()
	database.StartDatabase("test")
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

	accountOriginId, _ := serviceAccount.CreateAccount(accountOrigin)
	accountOrigin.ID = accountOriginId

	accountDestinationId, _ := serviceAccount.CreateAccount(accountDestination)
	accountDestination.ID = accountDestinationId

	transfer, _ := models.NewTransfer(accountOriginId, accountDestinationId, 500)

	servicesTransfer.CreateTransfer(transfer)

	t.Run("success", func(t *testing.T) {
		result, err := servicesTransfer.ListAllTransfersByAccount(accountOriginId)
		require.Nil(t, err)
		assert.NotEmpty(t, result[0].ID)
		assert.Len(t, result, 1)
	})

	t.Run("Error: Account not found", func(t *testing.T) {
		idAccountInvalid := uuid.NewV4().String()
		result, err := servicesTransfer.ListAllTransfersByAccount(idAccountInvalid)
		require.Error(t, err)
		require.EqualError(t, err, "Account not found")
		assert.Len(t, result, 0)
	})

}
