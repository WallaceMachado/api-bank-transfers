package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransfer(c *gin.Context) {

	result, _ := c.Get("accountId")
	c.JSON(http.StatusOK, result)
}
