package models

import (
	"gorm.io/gorm"
)

type Offer struct {
	ID        uint   `gorm:"primaryKey"`
	SellerID  int    `gorm:"type:int"`
	ProductID int    `gorm:"type:int"`
	Price     string `gorm:"type:decimal(10,2)"`
	Quantity  int    `gorm:"type:int"`
}

func (offer Offer) Migrate(databaseConnection *gorm.DB) error {
	return databaseConnection.AutoMigrate(offer)
}
