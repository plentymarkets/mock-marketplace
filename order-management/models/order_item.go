package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey"`
	OrderID  uint    `gorm:"type:int"`
	OfferID  uint    `gorm:"type:int"`
	GTIN     string  `gorm:"type:varchar(20)"`
	Price    float64 `gorm:"type:decimal(10,2)"`
	Quantity int     `gorm:"type:int"`
	SKU      string  `gorm:"type:varchar(255)"`
}

func (OrderItem OrderItem) Migrate(databaseConnection *gorm.DB) error {
	return databaseConnection.AutoMigrate(OrderItem)
}
