package main

import (
	"server/src/controllers"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router := app.Group("/api")

	router.GET("/city/:country", controllers.CityController)

	app.Run("0.0.0.0:" + os.Getenv("APP_PORT"))
}
