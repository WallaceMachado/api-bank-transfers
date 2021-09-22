package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/wallacemachado/api-bank-transfers/src/utils/validation"
)

type Login struct {
	Cpf    string `json:"cpf" valid:"notnull"`
	Secret string `json:"secret" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewLogin(cpf string, secret string) (*Login, error) {
	login := &Login{

		Cpf:    cpf,
		Secret: secret,
	}

	login.Cpf = validation.Format(login.Cpf)

	err := login.validate()

	if err != nil {
		return nil, err
	}

	return login, nil
}

func (login *Login) validate() error {

	if len(login.Secret) < 6 || len(login.Secret) > 12 {

		return errors.New("The secret must be between 6 and 12 characters.")
	}

	if len(login.Cpf) != 11 {

		return errors.New("Invalid CPF")
	}

	_, err := govalidator.ValidateStruct(login)

	if err != nil {
		return err
	}

	return nil
}
