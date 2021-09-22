package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
	"github.com/wallacemachado/api-bank-transfers/src/services"
)

func CreateAccount(c *gin.Context) {
	var account models.Account

	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err = account.Prepare(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	result, err := service.CreateAccount(account)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	resp := responses.ResponseCreateAccount{}
	resp.ID = result

	c.JSON(http.StatusCreated, resp)
}

func ListAllAccounts(c *gin.Context) {

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	result, err := service.ListAllAccounts()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetBalance(c *gin.Context) {
	id := c.Param("account_id")

	repository := &repositories.AccountRepository{}

	service := services.NewAccountService(repository)

	result, err := service.GetBalance(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := responses.ResponseGetBalance{}
	resp.Balance = result

	c.JSON(http.StatusOK, resp)
}
