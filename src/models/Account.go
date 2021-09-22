package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	utils "github.com/wallacemachado/api-bank-transfers/src/utils/security"
	"github.com/wallacemachado/api-bank-transfers/src/utils/validation"
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

func (account *Account) Prepare() error {

	if len(account.Secret) < 6 {
		fmt.Println(account.Secret)
		return errors.New("The secret must be at least 6 characters.")
	}

	secret, err := utils.Hash(account.Secret)

	if err != nil {
		return err
	}

	account.ID = uuid.NewV4().String()
	account.Secret = string(secret)

	account.Cpf = validation.Format(account.Cpf)

	err = account.validate()

	if err != nil {
		return err
	}

	return nil

}

func (account *Account) validate() error {

	err := validation.ValidateCPF(account.Cpf)
	if err != nil {
		return errors.New("invalid CPF")
	}
	_, err = govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}
