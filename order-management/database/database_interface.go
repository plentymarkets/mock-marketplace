package database

import "gorm.io/gorm"

type Database interface {
	SetupDatabase(dsn string) error
	GetConnection() *gorm.DB
}
