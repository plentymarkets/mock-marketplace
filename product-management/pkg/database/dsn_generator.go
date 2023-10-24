package database

import (
	"fmt"
	"os"
)

func GetMariaDBDSN() string {
	// Get the database connection details from environment variables
	dbHost := os.Getenv("MYSQL_TCP_HOST")          // Defaults to localhost
	dbPort := os.Getenv("PRODUCTS_MYSQL_TCP_PORT") // Defaults to 3306
	dbUser := os.Getenv("MYSQL_USER")              // Defaults to root
	dbPass := os.Getenv("MYSQL_PASSWORD")          // Defaults to empty string
	dbName := os.Getenv("MYSQL_DATABASE")

	// Create the database connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	return dsn
}
