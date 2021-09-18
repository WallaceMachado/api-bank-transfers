package responses

import "time"

type ResponseCreateAccount struct {
	ID uint `json:"id"`
}

type ResponseGetAccount struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	Balance   float32   `json:"balance,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
