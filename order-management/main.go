package main

import (
	"log"
	"order-management/database"
	"order-management/helper"
	"order-management/middlewares/authenticator"
	"order-management/migrate"
	"order-management/repositories"
	"order-management/routes"
	"order-management/seed"
	"os"
)

func init() {
	helper.LoadEnvVariables()
}

func main() {
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"), os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal("Could not create database")
	}
	databaseConnection := databaseFactory.GetConnection()

	authenticatorFactory, err := authenticator.CreateAuthenticator(os.Getenv("AUTHENTICATOR_DRIVER"))
	if err != nil {
		log.Fatal("Could not create authenticator")
	}
	authenticatorService := authenticatorFactory.NewAuthenticator(routes.GetExternalRoutesConfig()["authenticationService"])

	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)
	routes.RegisterRoutes(databaseConnection, authenticatorService)

	orderRepository := repositories.NewRepository(databaseConnection)
	orders := orderRepository.GetOrdersBySellerId(6944156350795021744)
	println(orders)
}
