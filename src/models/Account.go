package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"github.com/wallacemachado/api-bank-transfers/src/utils"
)

type Account struct {
	ID        string    `json:"id" gorm:"ype:uuid;primaryKeyt" valid:"notnull,uuid"`
	Name      string    `json:"name" valid:"notnull"`
	Cpf       string    `json:"cpf" gorm:"type:varchar(11);unique" valid:"notnull"`
	Secret    string    `json:"-" valid:"notnull"`
	Balance   float64   `json:"balance" valid:"notnull,float"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
}

func (Account *Account) Prepare() error {

	secret, err := utils.Hash(Account.Secret)

	if err != nil {
		return err
	}

	Account.ID = uuid.NewV4().String()
	Account.Secret = string(secret)

	err = Account.validate()

	if err != nil {
		return err
	}

	return nil

}

func (Account *Account) validate() error {

	_, err := govalidator.ValidateStruct(Account)

	if err != nil {
		return err
	}

	return nil
}
