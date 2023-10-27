package models

type Product struct {
	ID            uint      `gorm:"primarykey"`
	UserID        uint      `gorm:"<-" json:"user_id,omitempty"`
	Name          string    `gorm:"<-" json:"name" binding:"required"`
	Description   string    `gorm:"<-" json:"description" binding:"required"`
	GTIN          string    `gorm:"<-" json:"gtin,omitempty"`
	Categories    int       `gorm:"<-" json:"categories" binding:"required"`
	Manufacturers int       `gorm:"<-" json:"manufacturers" binding:"required"`
	Attributes    string    `gorm:"<-" json:"attributes" binding:"required"`
	Deleted       bool      `gorm:"<-" json:"deleted"` // TODO - If the deleted is required, the request fails with error
	Variants      []Variant `gorm:"<-" json:"variants"`
}
