package models_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/models"
)

func TestNewTransfer(t *testing.T) {
	account_origin_id := uuid.NewV4().String()
	account_destination_id := uuid.NewV4().String()
	amount := 500.00

	t.Run("NewTransfer: success", func(t *testing.T) {
		result, err := models.NewTransfer(account_origin_id, account_destination_id, amount)
		require.Nil(t, err)
		assert.NotEmpty(t, result.ID)
	})

	t.Run("Incorrect Account Origin Id: Error", func(t *testing.T) {
		result, err := models.NewTransfer("id_invalid", account_destination_id, amount)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("Accounts_ids equal: Origin account cannot be equal to desttination account", func(t *testing.T) {
		result, err := models.NewTransfer(account_origin_id, account_origin_id, amount)
		assert.EqualError(t, err, "Origin account cannot be equal to desttination account")
		assert.Nil(t, result)
	})

	t.Run("Incorrect Account Origin Id: Error", func(t *testing.T) {
		result, err := models.NewTransfer(account_origin_id, "id_invalid", amount)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("Incorrect amount: amount cannot be less than R$1", func(t *testing.T) {
		result, err := models.NewTransfer(account_origin_id, account_destination_id, 0)
		assert.EqualError(t, err, "The amount must be between R$1 and R$5000")
		assert.Nil(t, result)

	})

	t.Run("Incorrect amount: amount cannot be longer than R$5000", func(t *testing.T) {
		result, err := models.NewTransfer(account_origin_id, account_destination_id, 5001)
		assert.EqualError(t, err, "The amount must be between R$1 and R$5000")
		assert.Nil(t, result)

	})

}
