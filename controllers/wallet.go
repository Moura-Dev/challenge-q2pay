package controllers

import (
	"challenge-q2pay/db"
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetWalletByUserID
// @Summary Get Wallet by user id
// @ID GetWalletByUserID
// @Accept  json
// @Produce  json
// @Param   userID      path   int     true  "userID"
// @Success 200 {object} models.Wallet	"ok"
// @Failure 500 {string} string "error"
// @Router /wallet/{userID} [get]
func GetWallet(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	wallet, err := repository.GetWallet(idInt)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error getting wallet",
		})
	} else {
		ctx.JSON(200, gin.H{
			"wallet": wallet,
		})
	}
}

// @Summary Add Balance to user
// @Description Add Balance to User
// @Accept  json
// @Produce  json
// @Param   value  body models.Deposit true  "value"
// @Param   userID      path   int     true  "userID"
// @Failure 500 {string} string "Error"
// @Router /user/{userID}/deposit [post]
// @Success 201 {struct} Balance added
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
		ctx.JSON(201, gin.H{
			"message": "Balance added",
		})
		_ = db.CommitTransaction(tx)
	}
}
