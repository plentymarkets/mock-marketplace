package models

type Variant struct {
	ID         uint   `gorm:"primarykey" json:"id,omitempty"`
	ProductID  uint   `gorm:"<-" json:"product_id,omitempty"`
	Name       string `gorm:"<-" json:"name,omitempty"`
	GTIN       string `gorm:"<-" json:"gtin,omitempty"`
	Attributes int    `gorm:"<-" json:"attributes,omitempty"`
}
