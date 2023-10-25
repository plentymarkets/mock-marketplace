package main

import (
	"log"
	"order-management/database"
	"order-management/middlewares/authenticator"
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

	authenticatorFactory, err := authenticator.NewAuthenticator(os.Getenv("AUTHENTICATOR_DRIVER"))

	if err != nil {
		log.Fatal("Could not create authenticator")
	}

	authenticatorService := authenticatorFactory.NewAuthenticator(routes.GetExternalRoutesConfig()["authenticationService"])

	migrate.Migrate(databaseConnection)                             // Mode it to a separate main.go in cmd/migrate
	seed.Seed(databaseConnection)                                   // Mode it to a separate main.go in cmd/migrate
	routes.RegisterRoutes(databaseConnection, authenticatorService) // Why is the middleware injected in the Routing system?
}
