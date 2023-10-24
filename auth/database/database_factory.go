package database

import "fmt"

func CreateDatabase(driver, dataSourceName string) (Database, error) {
	var database Database
	var err error

	switch driver {
	case "mariadb":
		database = &MariaDBDatabase{}
		err = database.NewDatabase(dataSourceName)
	case "mysql":
		// instantiate MySQL and other databases similarly
	default:
		return nil, fmt.Errorf("unknown driver: %s", driver)
	}

	if err != nil {
		return nil, err
	}

	return database, nil
}
