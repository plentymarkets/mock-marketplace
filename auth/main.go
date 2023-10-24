package main

import (
	"auth/migrate"
	"auth/routes"
	"auth/seed"
)

func main() {
	databaseConnection := helper.GetDatabaseConnection()

	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)
	routes.RegisterRoutes(databaseConnection)
}
