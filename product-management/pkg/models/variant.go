package models

type Variant struct {
	ID         uint   `gorm:"primarykey" json:"id,omitempty"`
	ProductID  uint   `gorm:"<-" json:"product_id,omitempty"`
	UserID     uint   `gorm:"<-" json:"user_id,omitempty"`
	Name       string `gorm:"<-" json:"name,omitempty"`
	GTIN       string `gorm:"<-" json:"gtin,omitempty"`
	Attributes string `gorm:"<-" json:"attributes,omitempty"`
	Deleted    bool   `gorm:"<-" json:"deleted"` // TODO - If the deleted is required, the request fails with error
}
