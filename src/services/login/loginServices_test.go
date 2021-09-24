package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/config"
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	servicesAccount "github.com/wallacemachado/api-bank-transfers/src/services/account"
	servicesLogin "github.com/wallacemachado/api-bank-transfers/src/services/login"
)

func createTestDB() {
	config.Init()
	db := database.StartDatabase(config.DBName)
	db.Exec("DROP DATABASE IF EXISTS test;")
	db.Exec("CREATE DATABASE test;")

}

func TestLogin(t *testing.T) {
	createTestDB()
	database.StartDatabase("test")
	defer database.CloseConn()

	//cpf gerado aleatoriamente no site: https://www.4devs.com.br/gerador_de_cpf
	cpf := "71844473015"
	secret := "123456"

	account, _ := models.NewAccount("teste", cpf, secret, 1000)

	repositoryAccount := &repositories.AccountRepository{}

	serviceAccount := servicesAccount.NewAccountService(repositoryAccount)

	serviceAccount.CreateAccount(account)

	serviceLogin := servicesLogin.NewLoginService(repositoryAccount)

	t.Run("Success", func(t *testing.T) {
		newLogin, _ := models.NewLogin(cpf, secret)

		result, err := serviceLogin.Login(newLogin)
		require.Nil(t, err)
		assert.NotEmpty(t, result.Token)

	})

	t.Run("Error: Non-existent CPF", func(t *testing.T) {
		newLoginInvalidAccount, _ := models.NewLogin("12312312312", secret)
		result, err := serviceLogin.Login(newLoginInvalidAccount)
		require.Error(t, err)
		assert.Empty(t, result.Token)
		assert.EqualError(t, err, "Invalid CPF or secret")
	})

	t.Run("Error: invalid secret", func(t *testing.T) {
		newLoginInvalidSecret, _ := models.NewLogin(cpf, "123457")
		result, err := serviceLogin.Login(newLoginInvalidSecret)
		require.Error(t, err)
		assert.Empty(t, result.Token)
		assert.EqualError(t, err, "Invalid CPF or secret")
	})

}
