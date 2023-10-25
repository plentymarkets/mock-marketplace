package main

import (
	"auth/helper"
	"auth/migrate"
	"auth/routes"
	"auth/seed"
)

func main() {
	databaseConnection := helper.GetDatabaseConnection()

	migrate.Migrate(databaseConnection) // This should be in a separate main.go
	seed.Seed(databaseConnection)       // This should be in a separate main.go

	routes.RegisterRoutes(databaseConnection)
}
