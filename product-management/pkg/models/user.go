package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	UUID      string    `gorm:"index" json:"uuid" binding:"required"`
	Token     string    `gorm:"<-" json:"token" binding:"required"`
	Products  []Product `gorm:"<-" json:"products" binding:"required"`
	Variants  []Variant `gorm:"<-" json:"variants" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
