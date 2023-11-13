package main

import (
	"order-microservice/pkg/database"
	"order-microservice/pkg/routes/internal_router"
	"os"
)

func main() {
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		panic(err.Error())
	}

	databaseConnection := databaseFactory.GetConnection()

	router := internal_router.NewInternalRouter()
	router.RegisterRoutes(databaseConnection)
}
