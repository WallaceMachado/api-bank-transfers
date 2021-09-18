package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
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

	result, err := services.CreateAccount(account)
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

	result, err := services.ListAllAccounts()
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
	newid, err := strconv.Atoi(id)

	result, err := services.GetBalance(newid)
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
