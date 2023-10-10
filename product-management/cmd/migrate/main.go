package main

import (
	"log"
	"product-management/pkg/database"
	"product-management/pkg/models"
)

func main() {
	db := database.NewMariaDBDatabase()

	err := db.GetConnection().AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err2 := db.GetConnection().AutoMigrate(&models.Variant{})
	if err2 != nil {
		log.Fatal(err2.Error())
	}
}
