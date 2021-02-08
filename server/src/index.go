package main

import (
	"server/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

		controllers.Gogo()
	})

	app.Run("0.0.0.0:8080")
}
