package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func TestNewLogin(t *testing.T) {

	cpf := "451.340.780-84"
	secret := "12345g"

	t.Run("NewLogin: success", func(t *testing.T) {
		result, err := models.NewLogin(cpf, secret)
		require.Nil(t, err)
		assert.Equal(t, result.Cpf, "45134078084")
	})

	t.Run("Incorrect Secret: Invalid CPF or secret", func(t *testing.T) {
		result, err := models.NewLogin(cpf, "")
		assert.EqualError(t, err, "Invalid CPF or secret")
		assert.Nil(t, result)

	})

	t.Run("Incorrect CPF: Invalid CPF or secret", func(t *testing.T) {

		result, err := models.NewLogin("", secret)
		assert.EqualError(t, err, "Invalid CPF or secret")
		assert.Nil(t, result)

	})

}
