package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/controllers/responses"
	"github.com/wallacemachado/api-bank-transfers/src/models"
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
