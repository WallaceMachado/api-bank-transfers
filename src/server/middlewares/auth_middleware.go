package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/utils/authentication"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId, err := authentication.ValidateToken(c)

		if err != nil {
			c.AbortWithStatus(401)
		}

		c.Set("accountId", accountId)
		c.Next()
	}
}
