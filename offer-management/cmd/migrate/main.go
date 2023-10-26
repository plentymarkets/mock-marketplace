package main

import (
	"log"
	"offer-management/pkg/database"
	"offer-management/pkg/models"
)

func main() {
	dsn := database.GetMariaDBDSN()

	db := database.CreateConnection(dsn)

	err := db.AutoMigrate(&models.Offer{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
