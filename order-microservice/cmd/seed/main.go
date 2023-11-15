package main

import (
	"order-microservice/pkg/database"
	"order-microservice/pkg/seed"
	"order-microservice/pkg/utils/env-handler"
	"os"
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
