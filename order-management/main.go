package main

import (
	"order-management/helper"
	"order-management/migrate"
	"order-management/repositories"
	"order-management/routes"
	"order-management/seed"
)

func init() {
	helper.LoadEnvVariables()
}

func main() {
	databaseConnection := helper.GetDatabaseConnection()
	migrate.Migrate(databaseConnection)
	seed.Seed(databaseConnection)
	routes.RegisterRoutes(databaseConnection)

	orderRepository := repositories.NewRepository(databaseConnection)
	orders := orderRepository.GetOrdersBySellerId(6944156350795021744)
	println(orders)
}
