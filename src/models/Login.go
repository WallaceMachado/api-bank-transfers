package models

import "github.com/asaskevich/govalidator"

type Login struct {
	Cpf    string `json:"cpf" valid:"notnull"`
	Secret string `json:"secret" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (login *Login) Validate() error {

	_, err := govalidator.ValidateStruct(login)

	if err != nil {
		return err
	}

	return nil
}
