package controllers

import (
	"go_rest_gin/initializer"
	"go_rest_gin/models"
	"net/http"

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

func UpdateMessage(c *gin.Context) {
	// Get the value from request body and the id from the URL params
	var post models.Post
	id := c.Param("id")
	c.Bind(&post)

	// Get the original data from DB based on the id specified in the URL param
	var orig_post models.Post
	initializer.DB.First(&orig_post, id)

	// This will pass the orig_post and will be updated with the new value specified in post object.
	result := initializer.DB.Model(&orig_post).Updates(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the result.
	c.IndentedJSON(http.StatusOK, gin.H{
		"updated": orig_post,
	})
}

func DeleteMessage(c *gin.Context) {
	id := c.Param("id")

	// Delete from the database -> Actually it will only marked the deleted at column
	result := initializer.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		c.Status(http.StatusNotFound)

		return
	}
	log.Println("Row affected:", result.RowsAffected)
	// Respond
	c.JSON(200, gin.H{
		"post": "success",
	})
}
