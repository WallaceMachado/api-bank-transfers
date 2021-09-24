package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wallacemachado/api-bank-transfers/src/models"
	"github.com/wallacemachado/api-bank-transfers/src/repositories"
	"github.com/wallacemachado/api-bank-transfers/src/responses"
	services "github.com/wallacemachado/api-bank-transfers/src/services/transfer"
	"github.com/wallacemachado/api-bank-transfers/src/utils/dtos"
)

func CreateTransfer(c *gin.Context) {

	account_origin_id, _ := c.Get("accountId")
	var transfer dtos.CreateTransferDTO

	err := c.ShouldBindJSON(&transfer)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	transfer.Account_origin_id = account_origin_id.(string)

	newTransfer, err := models.NewTransfer(transfer.Account_origin_id, transfer.Account_destination_id, transfer.Amount)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repoAcc := &repositories.AccountRepository{}
	repoTransfer := &repositories.TransferRepository{}

	service := services.NewTransferService(repoAcc, repoTransfer)

	result, err := service.CreateTransfer(newTransfer)
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

	var transferResponse responses.ResponseTransfersByAccount

	for _, t := range result {
		if t.Account_origin_id == id {
			transferResponse.TranfersSent = append(transferResponse.TranfersSent, t)
		} else {

			transferResponse.TranfersReceived = append(transferResponse.TranfersReceived, t)
		}
	}

	c.JSON(http.StatusOK, transferResponse)
}
