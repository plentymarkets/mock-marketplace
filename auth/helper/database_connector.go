package helper

import (
	"auth/database"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetDatabaseConnection() *gorm.DB {
	databaseFactory, err := database.CreateDatabase(os.Getenv("DATABASE_DRIVER"), os.Getenv("MYSQL_DSN"))

	if err != nil {
		log.Fatal("Could not create database")
	}

	databaseConnection := databaseFactory.GetConnection()

	return databaseConnection
}
