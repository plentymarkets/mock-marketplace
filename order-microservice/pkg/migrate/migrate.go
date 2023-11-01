package migrate

import (
	"gorm.io/gorm"
	"order-microservice/pkg/models"
)

func Migrate(databaseConnection *gorm.DB) error {
	modelCollection := []Migrateable{
		models.Order{},
		models.OrderItem{},
	}

	for _, model := range modelCollection {
		err := model.Migrate(databaseConnection)
		if err != nil {
			return err
		}
	}

	return nil
}
