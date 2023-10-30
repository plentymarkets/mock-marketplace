package main

import (
	"order-management/database"
	"order-management/migrate"
	"order-management/routes"
	"order-management/seed"
	"os"
)

func main() {
	databaseFactory := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))
	databaseConnection := databaseFactory.GetConnection()

	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)

	var router routes.Router
	router = router.NewRouter(databaseConnection)
	router.RegisterRoutes()
}
