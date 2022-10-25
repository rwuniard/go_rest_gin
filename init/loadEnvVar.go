package init

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	// Loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
