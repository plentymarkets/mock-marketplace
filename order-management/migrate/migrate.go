package migrate

import (
	"gorm.io/gorm"
	"order-management/models"
)

func Migrate(databaseConnection *gorm.DB) {
	modelCollection := map[string]Migrateable{
		"order":      models.Order{},
		"order_item": models.OrderItem{},
	}

	for modelName, model := range modelCollection {
		err := model.Migrate(databaseConnection)
		if err != nil {
			panic("Could not migrate " + modelName + " model")
		}
	}
}
