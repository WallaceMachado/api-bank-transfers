package models

import "time"

type Account struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf" gorm:"unique_index"`
	Secret    string    `json:"-"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}
