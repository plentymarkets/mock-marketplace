package main

import (
	"os"
	"user-microservice/pkg/database"
	"user-microservice/pkg/seed"
	"user-microservice/pkg/utils/env-handler"
)

func main() {
	env_handler.LoadEnvironment()

	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		panic(err.Error())
	}

	databaseConnection := databaseFactory.GetConnection()

	err = seed.Seed(databaseConnection)

	if err != nil {
		panic(err.Error())
	}
}
