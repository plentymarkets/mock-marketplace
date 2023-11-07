package models

type Product struct {
	ID     uint   `gorm:"primarykey"`
	UserID uint   `gorm:"<-" json:"-"`
	Name   string `gorm:"<-" json:"name" binding:"required"`
	GTIN   string `gorm:"<-" json:"gtin" binding:"required"`
	SKU    string `gorm:"<-" json:"sku"`
	Offers []Offer
}
