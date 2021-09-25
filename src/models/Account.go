package models

import (
	"errors"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"github.com/wallacemachado/api-bank-transfers/src/shared/security"
	utils "github.com/wallacemachado/api-bank-transfers/src/shared/security"
	validation "github.com/wallacemachado/api-bank-transfers/src/shared/utils"
)

type Account struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey" valid:"notnull,uuid"`
	Name      string    `json:"name" valid:"notnull"`
	Cpf       string    `json:"cpf" gorm:"type:varchar(11);unique" valid:"notnull"`
	Secret    string    `json:"secret,omitempty" valid:"notnull"`
	Balance   float64   `json:"balance" valid:"notnull,float"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewAccount(name string, cpf string, secret string, balance float64) (*Account, error) {
	account := &Account{
		Name:    name,
		Cpf:     cpf,
		Secret:  secret,
		Balance: balance,
	}

	err := account.prepare()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) prepare() error {

	err := security.ValidateSecretString(account.Secret)
	if err != nil {
		return err
	}

	secret, err := utils.Hash(account.Secret)

	if err != nil {
		return err
	}

	account.ID = uuid.NewV4().String()
	account.Secret = string(secret)

	account.Name = strings.TrimSpace(account.Name)

	err = account.validate()

	if err != nil {
		return err
	}

	return nil

}

func (account *Account) validate() error {

	cpf, err := validation.ValidateCPF(account.Cpf)

	if err != nil {
		return errors.New("invalid CPF")
	}

	account.Cpf = cpf

	if len(account.Name) < 3 || len(account.Name) > 100 {

		return errors.New("The name must be between 3 and 100 characters.")
	}

	if account.Balance < 1 {

		return errors.New("Initial balance must be at least R$1")
	}

	_, err = govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}
