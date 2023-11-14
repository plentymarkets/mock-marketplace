package main

import (
	"product-management/pkg/database"
	"product-management/pkg/seed"
	"product-management/pkg/utils/env-handler"
)

func main() {
	env_handler.LoadEnvironment()

	dsn := database.GetMariaDBDSN()

	db := database.CreateConnection(dsn)

	err := seed.Seed(db)

	if err != nil {
		panic(err.Error())
	}
}
