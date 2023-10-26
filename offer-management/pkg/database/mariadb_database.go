package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func CreateConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
