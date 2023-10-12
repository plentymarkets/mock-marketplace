package models

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ProductID uint `gorm:"<-"`
}
