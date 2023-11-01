package database

import (
	"fmt"
)

func CreateDatabase(driver string) (Database, error) {
	var database Database
	var err error

	switch driver {
	case "mariadb":
		database = &MariaDBDatabase{}
		dataSourceName := GenerateDataSourceName()
		err = database.NewDatabase(dataSourceName)
	default:
		database = nil
		err = fmt.Errorf("unknown driver: %s", driver)
	}

	return database, err
}
