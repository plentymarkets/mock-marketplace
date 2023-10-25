package main

import (
	"auth/database"
	"auth/migrate"
	"auth/routes"
	"auth/seed"
	"log"
	"os"
)

func main() {
	dataSourceName := database.NewDsn(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_TCP_PORT"), os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_TIMEZONE"))
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"), dataSourceName)

	if err != nil {
		log.Fatal("Could not create database")
	}

	databaseConnection := databaseFactory.GetConnection()

	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)
	routes.RegisterRoutes(databaseConnection)
}
