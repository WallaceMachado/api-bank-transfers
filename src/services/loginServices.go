package services

import (
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories/interfaces"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
	"github.com/wallacemachado/api-bank-transfers/src/utils/authentication"
	"github.com/wallacemachado/api-bank-transfers/src/utils/security"
)

type LoginService struct {
	repository interfaces.IAccountRepository
}

func NewLoginService(repo interfaces.IAccountRepository) *LoginService {
	return &LoginService{
		repository: repo,
	}
}

func (s *LoginService) Login(login models.Login) (responses.ResponseLogin, error) {

	account, err := s.repository.GetAccountByCpf(login.Cpf)
	if err != nil {
		return responses.ResponseLogin{}, err
	}

	err = security.IsCorrectSecret(account.Secret, login.Secret)
	if err != nil {
		return responses.ResponseLogin{}, err
	}

	token, err := authentication.GenerateToken(account.ID)
	if err != nil {
		return responses.ResponseLogin{}, err
	}

	resp := responses.ResponseLogin{}
	resp.ID = account.ID
	resp.Name = account.Name
	resp.Token = token

	return resp, nil
}
