package main

import (
	"log"
	"product-management/pkg/database"
	"product-management/pkg/models"
)

func main() {
	dsn := database.GetMariaDBDSN()

	db := database.CreateConnection(dsn)

	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err2 := db.AutoMigrate(&models.Variant{})
	if err2 != nil {
		log.Fatal(err2.Error())
	}

	err3 := db.AutoMigrate(&models.User{})
	if err3 != nil {
		log.Fatal(err3.Error())
	}
}
