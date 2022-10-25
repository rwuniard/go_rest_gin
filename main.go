package main

import (
	"github.com/rwuniard/go_rest_gin/init"

	"github.com/gin-gonic/gin"
)

func init() {
	// Loading .env file
	init.LoadEnvVar()
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
