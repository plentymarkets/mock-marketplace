package helper

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}
}
