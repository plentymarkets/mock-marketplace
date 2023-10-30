package models

type Variant struct {
	ID         uint   `gorm:"primarykey" json:"-"`
	ProductID  uint   `gorm:"<-" json:"product_id,omitempty"`
	UserID     uint   `gorm:"<-" json:"user_id,omitempty"`
	Name       string `gorm:"<-" json:"name"`
	GTIN       string `gorm:"<-" json:"gtin"`
	Attributes string `gorm:"<-" json:"attributes"`
	Deleted    bool   `gorm:"<-" json:"deleted"`
}
