package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/wallacemachado/api-bank-transfers/src/config"
)

func GenerateToken(id string) (string, error) {
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claim["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.SecretKeyJwt))

}
