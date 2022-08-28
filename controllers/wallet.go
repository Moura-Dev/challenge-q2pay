package controllers

import (
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GET Wallet by user_id
func GetWallet(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	wallet, err := repository.GetWallet(idInt)
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
func DepositBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	balance := models.Transaction{}
	err = ctx.ShouldBindJSON(&balance)
	if err != nil {
		return
	}
	err = repository.AddBalance(idInt, balance.Value)
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
