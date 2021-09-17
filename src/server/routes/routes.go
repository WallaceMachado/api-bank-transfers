package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Config(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		accountRouter := main.Group("accounts")
		{
			accountRouter.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Get Accounts"})
			})
		}
	}

	return router
}
