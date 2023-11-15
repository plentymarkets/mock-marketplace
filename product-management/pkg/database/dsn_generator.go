package database

import (
	"fmt"
	"net/url"
	"os"
)

func GetMariaDBDSN() string {
	// Get the database connection details from environment variables
	dbHost := url.QueryEscape(os.Getenv("MYSQL_HOST"))
	dbPort := url.QueryEscape(os.Getenv("MYSQL_PORT"))
	dbUser := url.QueryEscape(os.Getenv("MYSQL_USER"))
	dbPass := url.QueryEscape(os.Getenv("MYSQL_PASSWORD"))
	dbName := url.QueryEscape(os.Getenv("MYSQL_DATABASE"))

	// Create the database connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	return dsn
}
