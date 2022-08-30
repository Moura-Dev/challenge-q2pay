package routes

import (
	"challenge-q2pay/controllers"
	_ "challenge-q2pay/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Basic API Transfer Balance
// @version 1.0
// @host localhost:5000
// @BasePath /api

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
			url := ginSwagger.URL("http://localhost:5000/api/docs/doc.json") // The url pointing to API definition
			routers.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		}
	}
	return router
}
