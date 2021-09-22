package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
	services "github.com/wallacemachado/api-bank-transfers/src/services/account"
	"github.com/wallacemachado/api-bank-transfers/src/utils/dtos"
)

func CreateAccount(c *gin.Context) {
	var createAccountDTO dtos.CreateAccountDTO

	err := c.ShouldBindJSON(&createAccountDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	account, err := models.NewAccount(createAccountDTO.Name, createAccountDTO.Cpf, createAccountDTO.Secret, createAccountDTO.Balance)

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
