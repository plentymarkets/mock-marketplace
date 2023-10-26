package database

import "gorm.io/gorm"

type Database interface {
	NewDatabase(dsn string) error
	GetConnection() *gorm.DB
}
