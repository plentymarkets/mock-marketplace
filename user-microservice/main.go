package main

import (
	"log"
	"os"
	"user-microservice/pkg/database"
	"user-microservice/pkg/routes"
	"user-microservice/pkg/utils/env-handler"
)

func main() {
	env_handler.LoadEnvironment()

	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		log.Fatal("Could not create database")
	}

	databaseConnection := databaseFactory.GetConnection()

	routes.RegisterRoutes(databaseConnection)
}
