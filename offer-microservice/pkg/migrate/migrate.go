package migrate

import (
	"gorm.io/gorm"
	"offer-microservice/pkg/models"
)

func Migrate(databaseConnection *gorm.DB) error {
	modelCollection := []Migrateable{
		models.offer{},
	}

	for _, model := range modelCollection {
		err := model.Migrate(databaseConnection)
		if err != nil {
			return err
		}
	}

	return nil
}
