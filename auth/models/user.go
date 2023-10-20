package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email           string    `gorm:"varchar(255)"`
	Password        string    `gorm:"varchar(255)"`
	Token           string    `gorm:"varchar(255)"`
	TokenExpiration time.Time `gorm:"timestamp"`
	roles           []Role
}

type Role struct {
	gorm.Model
	RoleName string `gorm:"varchar(255)"`
}
