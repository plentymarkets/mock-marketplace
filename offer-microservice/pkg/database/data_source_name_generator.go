package database

import (
	"fmt"
	"net/url"
	"os"
)

func GenerateDataSourceName() string {
	dbHost := url.QueryEscape(os.Getenv("MYSQL_HOST"))
	dbPort := url.QueryEscape(os.Getenv("MYSQL_PORT"))
	dbUser := url.QueryEscape(os.Getenv("MYSQL_USER"))
	dbPass := url.QueryEscape(os.Getenv("MYSQL_PASSWORD"))
	dbName := url.QueryEscape(os.Getenv("MYSQL_DATABASE"))
	dbTimezone := url.QueryEscape(os.Getenv("MYSQL_TIMEZONE"))

	dataSourceName := generate(dbUser, dbPass, dbHost, dbPort, dbName, dbTimezone)

	return dataSourceName
}

func generate(dbUser string, dbPass string, dbHost string, dbPort string, dbName string, dbTimezone string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbTimezone,
	)
}
