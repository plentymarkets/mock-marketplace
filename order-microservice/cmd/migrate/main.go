package main

import (
	"order-microservice/pkg/database"
	"order-microservice/pkg/migrate"
	"os"
)

func main() {
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		panic(err.Error())
	}

	databaseConnection := databaseFactory.GetConnection()

	err = migrate.Migrate(databaseConnection)

	if err != nil {
		panic(err.Error())
	}
}
