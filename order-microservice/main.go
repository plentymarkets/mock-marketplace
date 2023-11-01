package main

import (
	"order-microservice/pkg/database"
	"order-microservice/pkg/routes"
	"os"
)

func main() {
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"))

	if err != nil {
		panic(err.Error())
	}

	databaseConnection := databaseFactory.GetConnection()

	var router routes.Router
	router = router.NewRouter(databaseConnection)
	router.RegisterRoutes()
}
