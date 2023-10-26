package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDBDatabase struct {
	db *gorm.DB
}

func (db *MariaDBDatabase) NewDatabase(dsn string) error {
	var err error
	db.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func (db *MariaDBDatabase) GetConnection() *gorm.DB {
	return db.db
}
