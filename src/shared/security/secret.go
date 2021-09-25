package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

func ValidateSecretHash(secretWithHash, secretString string) error {
	return bcrypt.CompareHashAndPassword([]byte(secretWithHash), []byte(secretString))
}

func ValidateSecretString(secretString string) error {
	if len(secretString) < 6 || len(secretString) > 32 {

		return errors.New("The secret must be between 6 and 32 characters.")
	}

	return nil
}
