package repositories

import (
	"gorm.io/gorm"
	"order-management/models"
)

type OrderRepository struct {
	databaseConnection *gorm.DB
}

func SetupRepository(databaseConnection *gorm.DB) OrderRepository {
	repository := OrderRepository{}
	repository.databaseConnection = databaseConnection
	return repository
}

func (repository *OrderRepository) CreateOrder(order models.Order) {
	repository.databaseConnection.Create(&order)
}
