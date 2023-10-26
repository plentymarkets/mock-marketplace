package models

type Product struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"<-" json:"name" binding:"required"`
	SKU    string `gorm:"<-" json:"sku" binding:"required"`
	Offers []Offer
}
