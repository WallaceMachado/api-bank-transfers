package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/wallacemachado/api-bank-transfers/src/utils/security"
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

	err := login.validate()

	if err != nil {
		return nil, err
	}

	return login, nil
}

func (login *Login) validate() error {

	cpf, err := validation.ValidateCPF(login.Cpf)
	if err != nil {
		return errors.New("Invalid CPF or secret")
	}

	login.Cpf = cpf

	err = security.ValidateSecretString(login.Secret)
	if err != nil {
		return errors.New("Invalid CPF or secret")
	}

	_, err = govalidator.ValidateStruct(login)

	if err != nil {
		return err
	}

	return nil
}
