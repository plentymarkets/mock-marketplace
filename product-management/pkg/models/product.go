package models

type Product struct {
	ID            uint      `gorm:"primarykey" json:"id,omitempty"`
	Name          string    `gorm:"<-" json:"name,omitempty"`
	Description   string    `gorm:"<-" json:"description,omitempty"`
	Categories    int       `gorm:"<-" json:"categories,omitempty"`
	Manufacturers int       `gorm:"<-" json:"manufacturers,omitempty"`
	Deleted       bool      `gorm:"<-" json:"deleted,omitempty"`
	Attributes    int       `gorm:"<-" json:"attributes,omitempty"`
	Variants      []Variant `gorm:"<-" json:"variants,omitempty"`
}
