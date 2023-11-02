package main

import (
	"log"
	"offer-management/pkg/database"
	"offer-management/pkg/models"
)

func main() {
	dsn := database.GetMariaDBDSN() // Got Empty, want stuff

	db := database.CreateConnection(dsn) // Hotfix, insert .env file via play button

	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&models.Offer{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
