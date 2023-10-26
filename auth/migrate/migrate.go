package migrate

import (
	"auth/models"
	"gorm.io/gorm"
)

func Migrate(databaseConnection *gorm.DB) {
	err := databaseConnection.AutoMigrate(&models.User{})

	if err != nil {
		panic("Could not migrate database")
	}
}
