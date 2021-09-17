package models

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created"`
}
