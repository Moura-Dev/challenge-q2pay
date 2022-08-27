package controllers

import (
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"github.com/gin-gonic/gin"
	"strings"
)

func CreateTransaction(ctx *gin.Context) {
	transaction := models.Transaction{}
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		return
	}
	// set Transaction status upper case
	transaction.Status = strings.ToUpper(transaction.Status)
	err = repository.CreateTransaction(&transaction)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		ctx.JSON(200, gin.H{
			"message": "Transaction created",
		})
	}
}
