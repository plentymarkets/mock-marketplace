package migrate

import (
	"gorm.io/gorm"
	"user-microservice/pkg/models"
)

func Migrate(databaseConnection *gorm.DB) error {
	err := databaseConnection.AutoMigrate(&models.User{})

	if err != nil {
		return err
	}

	return nil
}
