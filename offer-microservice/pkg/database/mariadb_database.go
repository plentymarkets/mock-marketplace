package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDBDatabase struct {
	database *gorm.DB
}

func (database *MariaDBDatabase) NewDatabase(dataSourceName string) error {
	var err error
	database.database, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	return err
}

func (database *MariaDBDatabase) GetConnection() *gorm.DB {
	return database.database
}
