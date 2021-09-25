package responses

type ResponseLogin struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
