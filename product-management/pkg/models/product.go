package models

type Product struct {
	ID            uint      `gorm:"primarykey" json:"product_id"`
	UserID        uint      `gorm:"<-" json:"-"`
	Name          string    `gorm:"<-" json:"product_name"`
	Categories    int       `gorm:"<-" json:"categories" binding:"required"`
	Manufacturers int       `gorm:"<-" json:"manufacturers" binding:"required"`
	Deleted       bool      `gorm:"<-" json:"deleted"`
	Variants      []Variant `gorm:"<-" json:"variants" validate:"required"`
}
