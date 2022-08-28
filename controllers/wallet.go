package controllers

import (
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"github.com/gin-gonic/gin"
)

// GET Wallet by user_id
func GetWallet(ctx *gin.Context) {
	id := ctx.Param("id")
	wallet, err := repository.GetWallet(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"wallet": wallet,
		})
	}
}

// ADD balance user
func AddBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	balance := models.Transaction{}
	err := ctx.ShouldBindJSON(&balance)
	if err != nil {
		return
	}
	err = repository.AddBalance(id, balance.Value)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Balance added",
		})
	}
}
