package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Transfer struct {
	ID                     string    `json:"id" gorm:"type:uuid;primaryKey" valid:"notnull,uuid"`
	Account_origin_id      string    `json:"account_origin_id" valid:"notnull,uuid"`
	Account_origin         Account   `json:"-" valid:"-" gorm:"ForeignKey:Account_origin_id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Account_destination_id string    `json:"account_destination_id" valid:"notnull,uuid"`
	Account_destination    Account   `json:"-" valid:"-" gorm:"ForeignKey:Account_destination_id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Amount                 float64   `json:"amount" valid:"notnull,float"`
	CreatedAt              time.Time `json:"createdAt" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewTransfer(account_origin_id string, account_destination_idt string, amount float64) (*Transfer, error) {
	transfer := &Transfer{
		Account_origin_id:      account_origin_id,
		Account_destination_id: account_destination_idt,
		Amount:                 amount,
	}

	err := transfer.prepare()

	if err != nil {
		return nil, err
	}

	return transfer, nil
}

func (transfer *Transfer) prepare() error {
	transfer.ID = uuid.NewV4().String()

	err := transfer.validate()
	if err != nil {
		return err
	}

	return nil

}

func (transfer *Transfer) validate() error {

	if transfer.Amount < 1 || transfer.Amount > 5000 {
		return errors.New("The amount must be between R$1 and R$5000")
	}

	if transfer.Account_destination_id == transfer.Account_origin_id {

		return errors.New("Origin account cannot be equal to desttination account")
	}

	_, err := govalidator.ValidateStruct(transfer)

	if err != nil {
		return err
	}

	return nil
}
