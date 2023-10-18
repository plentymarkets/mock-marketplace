package repositories

import (
	"gorm.io/gorm"
	"order-management/models"
)

type OrderRepository struct {
	databaseConnection *gorm.DB
}

func NewRepository(databaseConnection *gorm.DB) OrderRepository {
	repository := OrderRepository{}
	repository.databaseConnection = databaseConnection
	return repository
}

func (repository *OrderRepository) CreateOrder(order models.Order) {
	repository.databaseConnection.Create(&order)
}

func (repository *OrderRepository) GetOrdersBySellerId(sellerId int) []models.Order {
	//CRITICAL: This is missing all kinds of error handling. If the Database has some problems for whatever reason we will not know about it. Implement it soon.

	var orders []models.Order
	repository.databaseConnection.Preload("OrderItems").Where("seller_id = ?", sellerId).Find(&orders)

	return orders
}
