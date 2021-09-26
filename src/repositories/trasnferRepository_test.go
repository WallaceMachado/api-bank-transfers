package repositories_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/shared/database"
)

var transfer = &models.Transfer{
	ID:                     uuid.NewV4().String(),
	Account_origin_id:      uuid.NewV4().String(),
	Account_destination_id: uuid.NewV4().String(),
	Amount:                 500,
	CreatedAt:              time.Now(),
}

func TestCreateTransfer(t *testing.T) {

	mock := NewMock()
	repo := &repositories.TransferRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.
			QuoteMeta(`INSERT INTO "transfers" ("id","account_origin_id","account_destination_id","amount","created_at") VALUES ($1,$2,$3,$4,$5)`)).
			WithArgs(
				transfer.ID,
				transfer.Account_origin_id,
				transfer.Account_destination_id,
				transfer.Amount,
				transfer.CreatedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		result, err := repo.CreateTransfer(transfer)

		assert.NotNil(t, result)
		assert.Equal(t, result.ID, transfer.ID)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		execError := errors.New("Exec Error")

		mock.ExpectBegin()
		mock.ExpectExec(regexp.
			QuoteMeta(`INSERT INTO "transfers" ("id","account_origin_id","account_destination_id","amount","created_at") VALUES ($1,$2,$3,$4,$5)`)).
			WithArgs(
				transfer.ID,
				transfer.Account_origin_id,
				transfer.Account_destination_id,
				transfer.Amount,
				transfer.CreatedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectCommit().WillReturnError(execError)

		result, err := repo.CreateTransfer(transfer)

		assert.Empty(t, result)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGetTransfersByAccountId(t *testing.T) {

	mock := NewMock()
	repo := &repositories.TransferRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery(regexp.
			QuoteMeta(`SELECT * FROM "transfers" WHERE account_origin_id = $1 OR Account_destination_id = $2`)).
			WithArgs(transfer.Account_origin_id, transfer.Account_origin_id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "account_origin_id", "account_destination_id", "amount", "created_at"}).
				AddRow(
					transfer.ID,
					transfer.Account_origin_id,
					transfer.Account_destination_id,
					transfer.Amount,
					transfer.CreatedAt))

		result, err := repo.GetTransfersByAccountId(transfer.Account_origin_id)

		assert.NotNil(t, result)
		assert.Equal(t, result[0].ID, transfer.ID)
		assert.Len(t, result, 1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		queryError := errors.New("Query Error")
		mock.ExpectQuery(regexp.
			QuoteMeta(`SELECT * FROM "transfers" WHERE account_origin_id = $1 OR Account_destination_id = $2`)).
			WithArgs(transfer.Account_origin_id, transfer.Account_origin_id).
			WillReturnError(queryError)

		result, err := repo.GetTransfersByAccountId(transfer.Account_origin_id)

		assert.Empty(t, result)
		assert.Len(t, result, 0)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

}
