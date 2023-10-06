package migrate

import (
	"gorm.io/gorm"
	"order-management/helper"
	"order-management/models"
	"reflect"
)

var databaseConnection *gorm.DB

func init() {
	helper.LoadEnvVariables()
	databaseConnection = helper.GetDatabaseConnection()
}

func Migrate() {
	modelCollection := []interface{}{
		&models.Order{},
		&models.OrderItem{},
	}

	for _, model := range modelCollection {
		err := databaseConnection.AutoMigrate(model)
		if err != nil {
			modelType := reflect.TypeOf(model).Elem().Name()
			panic("Could not migrate " + modelType + " model")
		}
	}
}
