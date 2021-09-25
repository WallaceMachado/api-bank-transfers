package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
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

func getToken(c *gin.Context) (string, error) {

	const Bearer_schema = "Bearer "
	header := c.GetHeader("Authorization")

	if header == "" {
		return "", errors.New("non-existent token!!")
	}

	token := header[len(Bearer_schema):]

	return token, nil
}

func ValidateToken(c *gin.Context) (string, error) {

	tokenString, err := getToken(c)

	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signature method %v", t.Header["alg"])
		}
		return []byte(config.SecretKeyJwt), nil
	})

	if err != nil {

		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountID := claims["id"].(string)

		return accountID, nil
	}

	return "", errors.New("invalid token")
}
