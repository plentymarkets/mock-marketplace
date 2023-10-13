package models

type Variant struct {
	ID         uint   `gorm:"primarykey"`
	ProductID  uint   `gorm:"<-"`
	Name       string `gorm:"<-"`
	GTIN       string `gorm:"<-"`
	Attributes int    `gorm:"<-"`
}
