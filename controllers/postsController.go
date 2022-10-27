package controllers

import (
	"go_rest_gin/initializer"
	"go_rest_gin/models"

	"log"

	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	// Get data from the request body
	var body models.Post

	// This will check if Content-Type == "application/json" or application/xml to determine how to parse input.
	c.Bind(&body)

	// Create a post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	// Write it to the database
	result := initializer.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(200, gin.H{
		"post": post,
	})
}

func GetMessagebyTitle(c *gin.Context) {

	var posts []models.Post

	title := c.Param("title")
	log.Println("The param is", title)

	result := initializer.DB.Where("title = ?", title).Find(&posts)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.IndentedJSON(200, gin.H{
		"post": posts,
	})
}
