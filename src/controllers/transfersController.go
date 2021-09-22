package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/services"
)

func CreateTransfer(c *gin.Context) {

	account_origin_id, _ := c.Get("accountId")
	var transfer models.Transfer

	err := c.ShouldBindJSON(&transfer)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	transfer.Account_origin_id = account_origin_id.(string)

	if err = transfer.Prepare(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repoAcc := &repositories.AccountRepository{}
	repoTransfer := &repositories.TransferRepository{}

	service := services.NewTransferService(repoAcc, repoTransfer)

	result, err := service.CreateTransfer(transfer)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, result)
}

func ListAllTransfersByAccount(c *gin.Context) {
	id, _ := c.Get("accountId")

	IdString := id.(string)

	repoAcc := &repositories.AccountRepository{}
	repoTransfer := &repositories.TransferRepository{}

	service := services.NewTransferService(repoAcc, repoTransfer)

	result, err := service.ListAllTransfersByAccount(IdString)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, result)
}
