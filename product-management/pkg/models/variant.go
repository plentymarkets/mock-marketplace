package models

type Variant struct {
	ID          uint   `gorm:"primarykey" json:"-"`
	ProductID   uint   `gorm:"<-" json:"product_id,omitempty"`
	GTIN        string `gorm:"<-" json:"gtin,omitempty"`
	Name        string `gorm:"<-" json:"variant_name"`
	Description string `gorm:"<-" json:"description" binding:"required"`
	Deleted     bool   `gorm:"<-" json:"deleted"`
}
