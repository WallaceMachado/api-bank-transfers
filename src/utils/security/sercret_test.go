package security_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wallacemachado/api-bank-transfers/src/utils/security"
)

func TestSecurity(t *testing.T) {

	secret := "123456"
	secretHash := ""
	t.Run("Hash: success", func(t *testing.T) {
		result, err := security.Hash(secret)
		secret = string(result)
		require.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("ValidateSecret: success", func(t *testing.T) {
		err := security.ValidateSecret(secretHash, secret)
		require.Error(t, err)

	})

	t.Run("ValidateSecret: error", func(t *testing.T) {
		err := security.ValidateSecret(secretHash, "secret invalid")
		require.Error(t, err)

	})

}
