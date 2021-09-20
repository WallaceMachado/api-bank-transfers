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
	Account_destination_id string    `json:"Account_destination_id" valid:"notnull,uuid"`
	Account_destination    Account   `json:"-" valid:"-" gorm:"ForeignKey:Account_destination_id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Amount                 float64   `json:"amount" valid:"notnull,float"`
	CreatedAt              time.Time `json:"createdAt" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (transfer *Transfer) Validate() error {

	if transfer.Amount < 1 || transfer.Amount > 5000 {
		return errors.New("The amount must be between 1 and 5000")
	}

	transfer.ID = uuid.NewV4().String()

	_, err := govalidator.ValidateStruct(transfer)

	if err != nil {
		return err
	}

	return nil
}
