package models

import (
	"time"
)

type Product struct {
	Id          int    `gorm:"primaryKey"`
	ProductName string `gorm:"<-"`
	Date        time.Time
	Deleted     uint `gorm:"<-"`
	Variants    []Variant
}
