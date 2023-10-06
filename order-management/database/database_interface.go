package database

import "gorm.io/gorm"

type DatabaseInterface interface {
	SetupDatabase(dsn string) error
	GetConnection() *gorm.DB
}
