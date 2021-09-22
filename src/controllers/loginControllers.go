package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/services"
)

func Login(c *gin.Context) {

	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	if err = login.Validate(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repository := &repositories.AccountRepository{}

	service := services.NewLoginService(repository)
	result, err := service.Login(login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)

}
