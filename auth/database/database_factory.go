package database

import (
	"fmt"
	"net/url"
)

func CreateDatabase(driver, dataSourceName string) (Database, error) {
	var database Database
	var err error

	switch driver {
	case "mariadb":
		database = &MariaDBDatabase{}
		dataSourceName = GetMariaDBDSN()
		err = database.NewDatabase(dataSourceName)
	default:
		database = nil
		err = fmt.Errorf("unknown driver: %s", driver)
	}

	return database, err
}

func NewDsn(dbUser string, dbPassword string, dbHost string, dbPort string, dbName string, dbTimezone string) string {
	dbTimezone = url.QueryEscape(dbTimezone)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbTimezone)
}
