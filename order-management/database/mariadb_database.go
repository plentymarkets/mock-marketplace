package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDBDatabase struct {
	db *gorm.DB
}

func (db *MariaDBDatabase) SetupDatabase(dataSourceName string) error {
	var err error
	db.db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	return err
}

func (db *MariaDBDatabase) GetConnection() *gorm.DB {
	return db.db
}
