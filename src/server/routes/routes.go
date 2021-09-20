package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/controllers"
	"github.com/wallacemachado/api-bank-transfers/src/server/middlewares"
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

		transferRouter := main.Group("transfers", middlewares.Auth())
		{
			transferRouter.POST("/", controllers.CreateTransfer)
			transferRouter.GET("/", controllers.ListAllTransfersByAccount)

		}
	}

	return router
}
