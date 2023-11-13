package main

import (
	"os"
	"user-microservice/pkg/database"
	"user-microservice/pkg/seed"
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
