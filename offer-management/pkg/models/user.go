package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primarykey"`
	UserName string `gorm:"<-" json:"user_name" binding:"required"`
	//UserPassword      string `gorm:"<-" json:"user_password" binding:"required"`
	EncryptedPassword string `gorm:"<-" json:"encrypted_password" binding:"required"`
	APIToken          string `gorm:"<-" json:"API_Token" binding:"required"`
	Deleted           bool   `gorm:"<-" json:"deleted"`
}

func (u *User) EncryptPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(hashedPassword)
	return nil
}

func (u *User) DecryptPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(providedPassword))
}
