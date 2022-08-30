package controllers

import (
	"challenge-q2pay/db"
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"challenge-q2pay/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Transfer Balance user to another user
// @Description Transfer Balance
// @Accept  json
// @Produce  json
// @Param   Transaction  body models.Transaction true  "value"
// @Failure 500 {string} string "Error"
// @Router /api/transfer [post]
// @Success 201 {string} string Balance transferred
func TransferBalance(ctx *gin.Context) {
	tx, err := db.StartTransaction()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	transaction := models.Transaction{}
	err = ctx.ShouldBindJSON(&transaction)
	if err != nil {
		return
	}
	userID := transaction.Payer
	//Get User by id
	user, err := repository.GetUser(userID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	//validate users or sellers
	if utils.ValidateCNPJ(user.CpfCnpj) {
		ctx.JSON(500, gin.H{
			"error": "Sellers can't transfer balance",
		})
		return
	}
	//Get wallet from payer
	wallet, err := repository.GetWallet(userID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error getting wallet from payer",
		})
		return
	}
	if transaction.Value.GreaterThan(wallet.Balance) {
		ctx.JSON(500, gin.H{
			"error": "Insufficient balance",
		})
		return
	}

	//Get wallet  from payee
	wallet, err = repository.GetWallet(transaction.Payee)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Error getting wallet from payee",
		})
		return
	}

	//Request authorization
	auth := utils.ResponseAuth{}
	AuthorizationResponse, err := utils.Authorization(utils.Url, &auth)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !AuthorizationResponse {
		ctx.JSON(500, gin.H{
			"error": "Authorization failed",
		})
		return
	}

	//Update balance from payer
	err = repository.RemoveBalance(transaction.Payer, transaction.Value, tx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Update balance from payee
	err = repository.AddBalance(transaction.Payee, transaction.Value, tx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		_ = db.RollbackTransaction(tx)
		return
	} else {
		_ = repository.CreateTransaction(&transaction)
		ctx.JSON(200, gin.H{
			"message": "Balance transferred",
		})
		_ = db.CommitTransaction(tx)
	}
}
