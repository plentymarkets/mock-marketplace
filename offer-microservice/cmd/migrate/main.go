package main

import (
	"offer-microservice/pkg/database"
	"offer-microservice/pkg/migrate"
	"offer-microservice/pkg/utils/env-handler"
	"os"
)

func main() {
	env_handler.LoadEnvironment()

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
