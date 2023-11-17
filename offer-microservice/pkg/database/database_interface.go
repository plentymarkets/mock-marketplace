package database

import "gorm.io/gorm"

type Database interface {
	NewDatabase(dataSourceName string) error
	GetConnection() *gorm.DB
}
