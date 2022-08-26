package controllers

import "github.com/gin-gonic/gin"

func HellowControllers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
