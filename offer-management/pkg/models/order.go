package models

type Order struct {
	ID        uint   `gorm:"primarykey"`
	OfferID   string `gorm:"<-" json:"offer_id" binding:"required"`
	UserID    string `gorm:"<-" json:"user_id" binding:"required"`
	Shipped   bool   `gorm:"<-" json:"shipped"`
	Delivered bool   `gorm:"<-" json:"delivered"`
	Deleted   bool   `gorm:"<-" json:"deleted"`
}
