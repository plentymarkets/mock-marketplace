package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"varchar(255)"`
	Password string `gorm:"varchar(255)"`
	Token    string `gorm:"varchar(255)"`
}
