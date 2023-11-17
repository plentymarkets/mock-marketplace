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

	dataSourceName := generate(dbUser, dbPass, dbHost, dbPort, dbName)

	return dataSourceName
}

func generate(dbUser string, dbPass string, dbHost string, dbPort string, dbName string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)
}
