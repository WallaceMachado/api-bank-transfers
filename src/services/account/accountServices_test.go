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
	services "github.com/wallacemachado/api-bank-transfers/src/services/account"
)

func createTestDB() {
	config.Init()
	db := database.StartDatabase(config.DBName)
	db.Exec("DROP DATABASE IF EXISTS test;")
	db.Exec("CREATE DATABASE test;")

}

func TestCreateAccount(t *testing.T) {
	createTestDB()
	database.StartDatabase("test")
	defer database.CloseConn()

	//cpf gerado aleatoriamente no site: https://www.4devs.com.br/gerador_de_cpf
	account, err := models.NewAccount("teste", "27714197005", "123456", 1000)

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	t.Run("Success", func(t *testing.T) {

		result, err := service.CreateAccount(account)
		require.Nil(t, err)
		assert.Equal(t, account.ID, result)
	})

	t.Run("Error: CPF already exists", func(t *testing.T) {

		_, err = service.CreateAccount(account)
		require.Error(t, err)
		assert.EqualError(t, err, "CPF already exists")
	})

}

func TestListAllAccounts(t *testing.T) {
	createTestDB()
	database.StartDatabase("test")
	defer database.CloseConn()

	//cpf gerado aleatoriamente no site: https://www.4devs.com.br/gerador_de_cpf
	account, _ := models.NewAccount("teste", "27714197005", "123456", 1000)

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	accountId, _ := service.CreateAccount(account)

	t.Run("Success", func(t *testing.T) {
		result, err := service.ListAllAccounts()
		require.Nil(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, account.ID, accountId)
	})

}

func TestGetBalance(t *testing.T) {
	createTestDB()
	database.StartDatabase("test")
	defer database.CloseConn()

	//cpf gerado aleatoriamente no site: https://www.4devs.com.br/gerador_de_cpf
	account, _ := models.NewAccount("teste", "27714197005", "123456", 1000)

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	accountId, _ := service.CreateAccount(account)

	t.Run("Success", func(t *testing.T) {
		result, err := service.GetBalance(accountId)
		require.Nil(t, err)
		assert.Equal(t, result, account.Balance)
	})

	t.Run("Error: Non-existent account", func(t *testing.T) {
		id := uuid.NewV4().String()
		_, err := service.GetBalance(id)
		require.Error(t, err)
		assert.EqualError(t, err, "Non-existent account")
	})

}
