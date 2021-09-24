package security

import "golang.org/x/crypto/bcrypt"

func Hash(secret string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
}

func ValidateSecret(secretWithHash, secretString string) error {
	return bcrypt.CompareHashAndPassword([]byte(secretWithHash), []byte(secretString))
}
