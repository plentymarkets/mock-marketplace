package main

import (
	"product-management/pkg/database"
	"product-management/pkg/seed"
)

func main() {
	dsn := database.GetMariaDBDSN()

	db := database.CreateConnection(dsn)

	err := seed.Seed(db)

	if err != nil {
		panic(err.Error())
	}
}
