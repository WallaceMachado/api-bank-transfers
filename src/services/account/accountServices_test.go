package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/database"
	"github.com/wallacemachado/api-bank-transfers/src/database/migrations"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	services "github.com/wallacemachado/api-bank-transfers/src/services/account"
)

func TestCreateAccount(t *testing.T) {

	database.StartDatabase("test")
	db := database.GetDatabase()
	defer database.CloseConn()

	account, err := models.NewAccount("teste", "27714197005", "123456", 1000)

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	result, err := service.CreateAccount(account)
	require.Nil(t, err)
	assert.Equal(t, account.ID, result)

	_, err = service.CreateAccount(account)
	require.Error(t, err)

	migrations.DeleteTablesTestDb(db)

}
