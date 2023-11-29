package main

import (
	"log"
	"product-management/pkg/database"
	"product-management/pkg/models"
	"product-management/pkg/utils/env-handler"
)

func main() {
	env_handler.LoadEnvironment()

	dsn := database.GetMariaDBDSN()

	db := database.CreateConnection(dsn)

	err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Variant{})

	if err != nil {
		log.Fatal(err.Error())
	}
}
