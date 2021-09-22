package dtos

type CreateAccountDTO struct {
	Name    string  `json:"name"`
	Cpf     string  `json:"cpf"`
	Secret  string  `json:"secret,omitempty"`
	Balance float64 `json:"balance"`
}
