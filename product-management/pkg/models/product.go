package models

type Product struct {
	ID            uint      `gorm:"primarykey" json:"-"`
	UserID        uint      `gorm:"<-" json:"-"`
	Categories    int       `gorm:"<-" json:"categories" binding:"required"`
	Manufacturers int       `gorm:"<-" json:"manufacturers" binding:"required"`
	Deleted       bool      `gorm:"<-" json:"deleted"`
	Variants      []Variant `gorm:"<-" json:"variants"`
}
