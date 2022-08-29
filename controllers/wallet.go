package controllers

import (
	"challenge-q2pay/db"
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
	tx, err := db.StartTransaction()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
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
	// Get Wallet by id
	_, err = repository.GetWallet(idInt)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error getting wallet to deposit",
		})
		return
	}
	err = repository.AddBalance(idInt, balance.Value, tx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		_ = db.RollbackTransaction(tx)
	} else {
		ctx.JSON(200, gin.H{
			"message": "Balance added",
		})
		_ = db.CommitTransaction(tx)
	}
}
