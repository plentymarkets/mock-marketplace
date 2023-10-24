package database

import "fmt"

func CreateDatabase(driver, dsn string) (Database, error) {
	var database Database
	var err error

	switch driver {
	case "mariadb":
		database = &MariaDBDatabase{}
		err = database.SetupDatabase(dsn)
	default:
		database = nil
		err = fmt.Errorf("unknown driver: %s", driver)
	}

	return database, err
}
