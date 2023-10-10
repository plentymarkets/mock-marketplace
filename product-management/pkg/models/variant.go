package models

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ProductId int `gorm:"<-"`
}
