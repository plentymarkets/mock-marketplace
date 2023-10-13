package models

type Product struct {
	ID            uint      `gorm:"primarykey"`
	Name          string    `gorm:"<-"`
	Description   string    `gorm:"<-"`
	Categories    int       `gorm:"<-"`
	Manufacturers int       `gorm:"<-"`
	Deleted       bool      `gorm:"<-"`
	Attributes    int       `gorm:"<-"`
	Variants      []Variant `gorm:"<-"`
}
