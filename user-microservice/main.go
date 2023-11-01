package main

import (
	"log"
	"os"
	"user-microservice/pkg/database"
	"user-microservice/pkg/routes"
)

func main() {
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		log.Fatal("Could not create database")
	}

	databaseConnection := databaseFactory.GetConnection()

	routes.RegisterRoutes(databaseConnection)
}
