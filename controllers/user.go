package controllers

import (
	"challenge-q2pay/models"
	"challenge-q2pay/repository"
	"challenge-q2pay/services"
	"challenge-q2pay/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Add a new user
// @Description Create User
// @Accept  json
// @Produce  json
// @Param   user      body models.User true  "user"
// @Failure 500 {string} web.APIError "Error"
// @Router /user/ [post]
// @Success 201 {object} models.User
func CreateUser(ctx *gin.Context) {
	user := models.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	//validate cpf or cnpj
	if !utils.ValidateCNPJ(user.CpfCnpj) && !utils.ValidateCPF(user.CpfCnpj) {
		ctx.JSON(500, gin.H{
			"error": "Invalid cpf or cnpj",
		})
		return
	}
	if !utils.ValidateEmail(user.Email) {
		ctx.JSON(500, gin.H{
			"error": "Invalid email",
		})
		return
	}
	user.CpfCnpj = utils.UnMaskCPFCNPJ(user.CpfCnpj)
	user.Password = services.SHA256ENCODER(user.Password)

	err = repository.CreateUser(&user)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(201, gin.H{
			"message": "User created",
		})
	}
}

// GetUserByID
// @Summary Get user by id
// @ID GetUserByID
// @Accept  json
// @Produce  json
// @Param   userID      path   int     true  "userID"
// @Success 200 {string} string	"ok"
// @Failure 500 {string} string "error"
// @Failure 404 {string} string "user not found"
// @Router /user/{userID} [get]
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	data, err := repository.GetUser(idInt)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}
	user := models.User{
		Id:       data.Id,
		FullName: data.FullName,
		CpfCnpj:  data.CpfCnpj,
		Email:    data.Email,
		Login:    data.Login,
		Active:   data.Active,
	}

	ctx.JSON(200, gin.H{
		"user": user,
	})
}
