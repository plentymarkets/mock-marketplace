package models

type BaseProduct struct {
	ID            uint      `gorm:"primarykey" json:"-"`
	UserID        uint      `gorm:"<-" json:"-"`
	Name          string    `gorm:"<-" json:"name" binding:"required"`
	Description   string    `gorm:"<-" json:"description" binding:"required"`
	GTIN          string    `gorm:"<-" json:"gtin,omitempty"`
	Categories    int       `gorm:"<-" json:"categories" binding:"required"`
	Manufacturers int       `gorm:"<-" json:"manufacturers" binding:"required"`
	Attributes    string    `gorm:"<-" json:"attributes" binding:"required"`
	Deleted       bool      `gorm:"<-" json:"deleted"`
	Variants      []Variant `gorm:"<-" json:"variants"`
}

type BaseVariant struct {
	BaseProduct
	Manufacturers int    `gorm:"<-" json:"manufacturers" binding:"required"`
	Attributes    string `gorm:"<-" json:"attributes" binding:"required"`
	Deleted       bool   `gorm:"<-" json:"deleted"`
}
