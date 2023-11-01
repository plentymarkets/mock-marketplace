package main

import (
	"order-microservice/pkg/database"
	"order-microservice/pkg/seed"
	"os"
)

func main() {
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
