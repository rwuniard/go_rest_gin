package main

import (
	"go_rest_gin/initializer"
	"go_rest_gin/models"
)

// Initializing the Environment variable and Database connection.
func init() {
	initializer.LoadEnvVar()
	initializer.ConnectPostgres()
}

// This is to initialize the schema defined in the models.Post struct.
func main() {
	initializer.DB.AutoMigrate(&models.Post{})
}
