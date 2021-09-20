package models

import (
	"time"
)

type Transfer struct {
	ID                     string    `json:"id" gorm:"type:uuid;primaryKey" valid:"notnull,uuid"`
	Account_origin_id      string    `json:"account_origin_id" valid:"notnull"`
	Account_origin         Account   `json:"-" valid:"-" gorm:"ForeignKey:Account_origin_id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Account_destination_id string    `json:"Account_destination_id" valid:"notnull"`
	Account_destination    Account   `json:"-" valid:"-" gorm:"ForeignKey:Account_destination_id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Amount                 float64   `json:"balance" valid:"notnull,float"`
	CreatedAt              time.Time `json:"createdAt" valid:"-"`
}
