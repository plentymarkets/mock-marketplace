package models

import (
	"time"
)

type User struct {
	ID              int       `gorm:"primaryKey"`
	Email           string    `gorm:"varchar(255)"`
	Password        string    `gorm:"varchar(255)"`
	Token           string    `gorm:"varchar(255)"`
	TokenExpiration time.Time `gorm:"timestamp"`
}
