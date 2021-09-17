package models

import "time"

type Account struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	Secret    string    `json:"secret,omitempty"`
	Balance   float32   `json:"balance,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
}
