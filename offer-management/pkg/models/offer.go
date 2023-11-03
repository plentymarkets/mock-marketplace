package models

type Offer struct {
	ID        uint    `gorm:"primarykey"`
	ProductID uint    `gorm:"<-" json:"product_id" binding:"required"`
	SKU       string  `gorm:"-:all" json:"sku" binding:"required"`
	Price     float32 `gorm:"<-" json:"price" binding:"required"`
	Currency  string  `gorm:"<-" json:"currency" binding:"required"`
	Quantity  int     `gorm:"<-" json:"quantity" binding:"required"`
	Discount  int     `gorm:"<-" json:"discount" binding:"required"`
	Deleted   bool    `gorm:"<-" json:"deleted"`
}
