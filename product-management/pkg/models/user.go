package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string    `gorm:"index" json:"uuid" binding:"required"`
	Token    string    `gorm:"<-" json:"token" binding:"required"`
	Products []Product `gorm:"<-" json:"products" binding:"required"`
	Variants []Variant `gorm:"<-" json:"variants" binding:"required"`
}
