package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/utils"
)

func TestNewAccount(t *testing.T) {
	name := "teste"
	cpf := "451.340.780-84"
	secret := "12345g"
	balance := 1000.00

	t.Run("NewAccount: success", func(t *testing.T) {
		result, err := models.NewAccount(name, cpf, secret, balance)
		require.Nil(t, err)
		assert.NotEmpty(t, result.ID)
		assert.Equal(t, result.Cpf, "45134078084")
	})

	t.Run("Incorrect Secret: password cannot be less than 6 characters", func(t *testing.T) {
		result, err := models.NewAccount(name, cpf, "12345", balance)
		assert.EqualError(t, err, "The secret must be between 6 and 32 characters.")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Secret: password cannot be longer than 32 characters", func(t *testing.T) {
		incorrectSecret := utils.GenerateString(33)
		result, err := models.NewAccount(name, cpf, incorrectSecret, balance)
		assert.EqualError(t, err, "The secret must be between 6 and 32 characters.")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Name: name cannot be less than 3 characters", func(t *testing.T) {
		result, err := models.NewAccount(" n   ", cpf, secret, balance)
		assert.EqualError(t, err, "The name must be between 3 and 100 characters.")
		assert.Nil(t, result)

	})

	t.Run("Incorrect Name: name cannot be longer than 100 characters", func(t *testing.T) {
		incorrectName := utils.GenerateString(101)
		result, err := models.NewAccount(incorrectName, cpf, secret, balance)
		assert.EqualError(t, err, "The name must be between 3 and 100 characters.")
		assert.Nil(t, result)

	})

	t.Run("Incorrect CPF: cpf cannot be null", func(t *testing.T) {

		result, err := models.NewAccount(name, "", secret, balance)
		assert.EqualError(t, err, "invalid CPF")
		assert.Nil(t, result)

	})

	t.Run("Incorrect balance: Initial balance must be at least R$1", func(t *testing.T) {

		result, err := models.NewAccount(name, cpf, secret, 0)
		assert.EqualError(t, err, "Initial balance must be at least R$1")
		assert.Nil(t, result)

	})

}
