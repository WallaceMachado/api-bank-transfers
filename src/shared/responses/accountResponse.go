package responses

type ResponseCreateAccount struct {
	ID string `json:"id"`
}

type ResponseGetBalance struct {
	Balance float64 `json:"balance"`
}
