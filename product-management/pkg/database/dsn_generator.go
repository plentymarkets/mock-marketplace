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
	dbTimezone := url.QueryEscape(os.Getenv("MYSQL_TIMEZONE"))

	// Create the database connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbTimezone,
	)

	return dsn
}
