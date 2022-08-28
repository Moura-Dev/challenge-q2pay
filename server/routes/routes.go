package routes

import (
	"challenge-q2pay/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		routers := main.Group("/")
		{
			routers.POST("/user", controllers.CreateUser)
			routers.GET("/user/:id", controllers.GetUser)
			routers.GET("/wallet/:id", controllers.GetWallet)
			//Deposit balance user
			routers.POST("/user/:id/deposit", controllers.DepositBalance)
			routers.POST("/transfer", controllers.TransferBalance)

		}
	}
	return router
}
