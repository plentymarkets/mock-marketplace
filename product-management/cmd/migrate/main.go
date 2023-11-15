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

	err1 := db.AutoMigrate(&models.User{})
	if err1 != nil {
		log.Fatal(err1.Error())
	}

	err2 := db.AutoMigrate(&models.Product{})
	if err2 != nil {
		log.Fatal(err2.Error())
	}

	err3 := db.AutoMigrate(&models.Variant{})
	if err3 != nil {
		log.Fatal(err3.Error())
	}
}
