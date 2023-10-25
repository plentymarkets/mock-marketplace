package models

type Product struct {
	ID            uint      `gorm:"primarykey"`
	UserID        uint      `gorm:"<-" json:"user_id,omitempty"`
	Name          string    `gorm:"<-" json:"name" binding:"required"`
	Description   string    `gorm:"<-" json:"description" binding:"required"`
	Categories    int       `gorm:"<-" json:"categories" binding:"required"`
	Manufacturers int       `gorm:"<-" json:"manufacturers" binding:"required"`
	Deleted       bool      `gorm:"<-" json:"deleted"` // TODO - If the deleted is required, the request fails with error
	Attributes    int       `gorm:"<-" json:"attributes" binding:"required"`
	Variants      []Variant `gorm:"<-" json:"variants"`
}
