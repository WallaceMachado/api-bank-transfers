package responses

import "time"

type ResponseCreateAccount struct {
	ID uint `json:"id"`
}

type ResponseGetAccount struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ResponseGetBalance struct {
	Balance float32 `json:"balance"`
}
