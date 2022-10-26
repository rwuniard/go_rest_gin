package controllers

import (
	"go_rest_gin/initializer"
	"go_rest_gin/models"

	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {

	post := models.Post{
		Title: "Hello",
		Body:  "This is the body",
	}

	result := initializer.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
