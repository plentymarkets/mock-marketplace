package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"user-microservice/pkg/database"
	"user-microservice/pkg/routes"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env.dev file")
	}

	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		log.Fatal("Could not create database")
	}

	databaseConnection := databaseFactory.GetConnection()

	routes.RegisterRoutes(databaseConnection)
}
