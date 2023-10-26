package models

type User struct {
	ID           uint   `gorm:"primarykey"`
	UserName     string `gorm:"<-" json:"user_name" binding:"required"`
	UserPassword string `gorm:"<-" json:"user_password" binding:"required"`
	APIToken     string `gorm:"<-" json:"API_Token" binding:"required"`
	Deleted      bool   `gorm:"<-" json:"deleted"` // TODO - If the deleted is required, the request fails with error
}
