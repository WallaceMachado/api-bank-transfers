package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/shared/authentication"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId, err := authentication.ValidateToken(c)

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)

		}

		c.Set("accountId", accountId)
		c.Next()
	}
}
