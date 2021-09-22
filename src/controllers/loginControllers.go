package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	services "github.com/wallacemachado/api-bank-transfers/src/services/login"
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

	newLogin, err := models.NewLogin(login.Cpf, login.Secret)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})

		return
	}

	repository := &repositories.AccountRepository{}

	service := services.NewLoginService(repository)
	result, err := service.Login(newLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)

}
