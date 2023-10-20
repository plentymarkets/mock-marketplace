package database

import "gorm.io/gorm"

type DatabaseInterface interface {
	NewDatabase(dsn string) error
	GetConnection() *gorm.DB
}
