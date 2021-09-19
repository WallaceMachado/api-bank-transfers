package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/controllers"
)

func Config(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		accountRouter := main.Group("accounts")
		{
			accountRouter.POST("/", controllers.CreateAccount)
			accountRouter.GET("/", controllers.ListAllAccounts)
			accountRouter.GET("/:account_id/balance", controllers.GetBalance)

		}

		loginRouter := main.Group("login")
		{
			loginRouter.POST("/", controllers.Login)

		}
	}

	return router
}
