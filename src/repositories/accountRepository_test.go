package repositories_test

import (
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/shared/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMock() sqlmock.Sqlmock {

	dbSqlMock, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSqlMock,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	database.SetDatabase(gormDB)

	return mock
}

var account = &models.Account{
	ID:        uuid.NewV4().String(),
	Name:      "teste",
	Cpf:       "09508813091",
	Secret:    "123456",
	Balance:   1000,
	CreatedAt: time.Now(),
}

func TestCreateAccount(t *testing.T) {

	mock := NewMock()
	repo := &repositories.AccountRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.
			QuoteMeta("INSERT INTO \"accounts\" (\"id\",\"name\",\"cpf\",\"secret\",\"balance\",\"created_at\") VALUES ($1,$2,$3,$4,$5,$6)")).
			WithArgs(
				account.ID,
				account.Name,
				account.Cpf,
				account.Secret,
				account.Balance,
				account.CreatedAt).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		id, err := repo.CreateAccount(account)

		assert.NotNil(t, id)
		assert.Equal(t, id, account.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		execError := errors.New("Exec Error")

		mock.ExpectBegin()
		mock.ExpectExec(regexp.
			QuoteMeta("INSERT INTO \"accounts\" (\"id\",\"name\",\"cpf\",\"secret\",\"balance\",\"created_at\") VALUES ($1,$2,$3,$4,$5,$6)")).
			WithArgs(
				account.ID,
				account.Name,
				account.Cpf,
				account.Secret,
				account.Balance,
				account.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectCommit().WillReturnError(execError)

		id, err := repo.CreateAccount(account)

		assert.Empty(t, id)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

}

func TestGetAllAccounts(t *testing.T) {

	mock := NewMock()
	repo := &repositories.AccountRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery(regexp.
			QuoteMeta("SELECT id,name,cpf,balance,created_at FROM \"accounts\"")).
			WithArgs().
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cpf", "balance", "created_at"}).
				AddRow(account.ID,
					account.Name,
					account.Cpf,
					account.Balance,
					account.CreatedAt))

		accounts, err := repo.GetAllAccounts()

		assert.NotNil(t, accounts)
		assert.Equal(t, accounts[0].ID, account.ID)
		assert.Len(t, accounts, 1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		queryError := errors.New("Query Error")
		mock.ExpectQuery(regexp.
			QuoteMeta("SELECT id,name,cpf,balance,created_at FROM \"accounts\"")).
			WithArgs().
			WillReturnError(queryError)

		accounts, err := repo.GetAllAccounts()

		assert.Nil(t, accounts)
		assert.Len(t, accounts, 0)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

}

func TestGetAccountById(t *testing.T) {

	mock := NewMock()
	repo := &repositories.AccountRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery(regexp.
			QuoteMeta(`SELECT * FROM "accounts" WHERE id =$1`)).
			WithArgs(account.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cpf", "secret", "balance", "created_at"}).
				AddRow(account.ID,
					account.Name,
					account.Cpf,
					account.Secret,
					account.Balance,
					account.CreatedAt))

		result, err := repo.GetAccountById(account.ID)

		assert.NotNil(t, result)
		assert.Equal(t, result.ID, result.ID)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		queryError := errors.New("Query Error")
		mock.ExpectQuery(regexp.
			QuoteMeta(`SELECT * FROM "accounts" WHERE id =$1`)).
			WithArgs(account.ID).
			WillReturnError(queryError)

		result, err := repo.GetAccountById(account.ID)

		assert.Empty(t, result.ID)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGetAccountByCPF(t *testing.T) {

	mock := NewMock()
	repo := &repositories.AccountRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {

		mock.ExpectQuery(regexp.
			QuoteMeta(`SELECT * FROM "accounts" WHERE cpf =$1`)).
			WithArgs(account.Cpf).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "cpf", "secret", "balance", "created_at"}).
				AddRow(account.ID,
					account.Name,
					account.Cpf,
					account.Secret,
					account.Balance,
					account.CreatedAt))

		result, err := repo.GetAccountByCpf(account.Cpf)

		assert.NotNil(t, result)
		assert.Equal(t, result.ID, account.ID)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		queryError := errors.New("Query Error")
		mock.ExpectQuery(regexp.
			QuoteMeta(`SELECT * FROM "accounts" WHERE cpf =$1`)).
			WithArgs(account.Cpf).
			WillReturnError(queryError)

		result, err := repo.GetAccountByCpf(account.Cpf)

		assert.Empty(t, result.ID)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestUpdateAccount(t *testing.T) {

	mock := NewMock()
	repo := &repositories.AccountRepository{}
	defer database.CloseConn()

	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.
			QuoteMeta(`UPDATE "accounts" SET "name"=$1,"cpf"=$2,"secret"=$3,"balance"=$4,"created_at"=$5 WHERE "id" = $6`)).
			WithArgs(account.Name, account.Cpf, account.Secret, account.Balance, account.CreatedAt, account.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		result, err := repo.UpdateAccount(*account)

		assert.NotNil(t, result)
		assert.Equal(t, result.ID, account.ID)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error", func(t *testing.T) {
		execError := errors.New("Exec Error")
		mock.ExpectBegin()
		mock.ExpectExec(regexp.
			QuoteMeta(`UPDATE "accounts" SET "name"=$1,"cpf"=$2,"secret"=$3,"balance"=$4,"created_at"=$5 WHERE "id" = $6`)).
			WithArgs(account.Name, account.Cpf, account.Secret, account.Balance, account.CreatedAt, account.ID).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit().WillReturnError(execError)

		result, err := repo.UpdateAccount(*account)

		assert.Empty(t, result.ID)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
