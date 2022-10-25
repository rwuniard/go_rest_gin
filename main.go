package main

import (
	"go_rest_gin/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	// Loading .env file
	initializer.LoadEnvVar()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
