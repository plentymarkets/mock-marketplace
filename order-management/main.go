package main

import (
	"log"
	"order-management/database"
	"order-management/migrate"
	"order-management/routes"
	"order-management/seed"
	"os"
)

func main() {
	dataSourceName := database.NewDsn(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_TCP_PORT"), os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_TIMEZONE"))
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"), dataSourceName)

	if err != nil {
		log.Fatal("Could not create database")
	}

	databaseConnection := databaseFactory.GetConnection()

	var externalRouter routes.ExternalRouter
	externalRouter = externalRouter.NewExternalRouter()

	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)
	routes.RegisterRoutes(databaseConnection, externalRouter)
}
