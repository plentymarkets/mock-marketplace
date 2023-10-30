package database

import (
	"fmt"
	"log"
	"net/url"
)

func CreateDatabase(driver string) Database {
	var database Database
	var err error

	switch driver {
	case "mariadb":
		database = &MariaDBDatabase{}
		dataSourceName := GetMariaDBDSN()
		err = database.SetupDatabase(dataSourceName)
	default:
		database = nil
		err = fmt.Errorf("unknown driver: %s", driver)
	}

	if err != nil {
		log.Fatal("Could not create database")
	}

	return database
}

func NewDsn(dbUser string, dbPassword string, dbHost string, dbPort string, dbName string, dbTimezone string) string {
	dbTimezone = url.QueryEscape(dbTimezone)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbTimezone)
}
