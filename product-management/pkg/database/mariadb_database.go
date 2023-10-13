package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type MariaDBDatabase struct {
	db *gorm.DB
}

func NewMariaDBDatabase() MariaDBDatabase {
	return MariaDBDatabase{
		db: createConnection(),
	}
}

func (db *MariaDBDatabase) GetConnection() *gorm.DB {
	return db.db
}

func createConnection() *gorm.DB {

	// Get the database connection details from environment variables
	dbHost := os.Getenv("MYSQL_TCP_HOST") // Defaults to localhost
	dbPort := os.Getenv("MYSQL_TCP_PORT") // Defaults to 3306
	dbUser := os.Getenv("MYSQL_USER")     // Defaults to root
	dbPass := os.Getenv("MYSQL_PASSWORD") // Defaults to empty string
	dbName := os.Getenv("MYSQL_DATABASE") // Defaults to database

	// Create the database connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
