package main

import (
	"go_rest_gin/controllers"
	"go_rest_gin/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	// Loading .env file
	initializer.LoadEnvVar()

	// Load DB connection
	initializer.ConnectPostgres()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreateMessage)
	r.GET("/posts/:title", controllers.GetMessagebyTitle)
	r.Run()
}
