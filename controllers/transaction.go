package controllers

import (
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"challenge-q2pay/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type responseAuth struct {
	Authorization bool
}

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

// Transfer balance from one account to another
func TransferBalance(ctx *gin.Context) {
	//db.StartDB()
	transaction := models.Transaction{}
	err := ctx.ShouldBindJSON(&transaction)
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
			"error": err.Error(),
		})
		return
	}
	if transaction.Value.GreaterThan(wallet.Balance) {
		ctx.JSON(500, gin.H{
			"error": "Insufficient balance",
		})
		return
	}
	//Request authorization
	auth := responseAuth{}
	auth.Authorization = utils.Authorization(utils.Url, auth)
	if !auth.Authorization {
		ctx.JSON(500, gin.H{
			"error": "Authorization failed",
		})
		return
	}

	//Update balance from payer
	err = repository.UpdateBalance(transaction.Payer, transaction.Value)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Get wallet  from payee
	wallet, err = repository.GetWallet(transaction.Payee)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Update balance from payee
	err = repository.AddBalance(transaction.Payee, transaction.Value)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"message": "Balance transferred",
		})
	}
}
