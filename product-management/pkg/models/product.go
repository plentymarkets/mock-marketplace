package models

type Product struct {
	ID            uint      `gorm:"primarykey"`
	Name          string    `gorm:"<-" json:"name"`
	Description   string    `gorm:"<-" json:"description"`
	Categories    int       `gorm:"<-" json:"categories"`
	Manufacturers int       `gorm:"<-" json:"manufacturers"`
	Deleted       bool      `gorm:"<-" json:"deleted"`
	Attributes    int       `gorm:"<-" json:"attributes"`
	Variants      []Variant `gorm:"<-" json:"variants"`
}
