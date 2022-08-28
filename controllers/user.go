package controllers

import (
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"challenge-q2pay/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	user := models.User{}
	ctx.ShouldBindJSON(&user)
	//validate cpf or cnpj
	if !utils.ValidateCNPJ(user.CpfCnpj) && !utils.ValidateCPF(user.CpfCnpj) {
		ctx.JSON(500, gin.H{
			"error": "Invalid cpf or cnpj",
		})
		return
	}
	err := repository.CreateUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "User created",
		})
	}
}

// Get User By id
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	data, err := repository.GetUser(idInt)
	user := models.User{
		Id:       data.Id,
		FullName: data.FullName,
		CpfCnpj:  data.CpfCnpj,
		Email:    data.Email,
		Login:    data.Login,
		Active:   data.Active,
	}

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"user": user,
		})
	}
}
