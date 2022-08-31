package routes

import (
	"challenge-q2pay/controllers"
	_ "challenge-q2pay/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title Basic API Transfer Balance
// @version 1.0
// @host localhost:PORT
// @BasePath /api

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	url := fmt.Sprintf("http://localhost:%s/api/docs/doc.json", os.Getenv("PORT"))
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
			url := ginSwagger.URL(url) // The url pointing to API definition
			routers.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		}
	}
	return router
}
