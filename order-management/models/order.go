package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID              uint      `gorm:"primaryKey"`
	CustomerID      int       `gorm:"type:int"`
	SellerID        int       `gorm:"type:int"`
	OrderNumber     string    `gorm:"type:varchar(255)"`
	OrderDate       time.Time `gorm:"type:datetime"`
	Status          string    `gorm:"type:varchar(50)"`
	TotalSum        float64   `gorm:"type:decimal(10,2)"`
	BillingAddress  string    `gorm:"type:varchar(255)"`
	ShippingAddress string    `gorm:"type:varchar(255)"`
	OrderItems      []OrderItem
}

func (Order Order) Migrate(databaseConnection *gorm.DB) error {
	return databaseConnection.AutoMigrate(Order)
}
