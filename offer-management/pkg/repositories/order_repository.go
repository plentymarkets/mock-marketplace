package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"offer-management/pkg/models"
)

type OrderRepository struct {
	database *gorm.DB
}

func NewOrderRepository(gormDB *gorm.DB) (*OrderRepository, error) {
	// if the database is nil, it will crash. Throw an error
	if gormDB == nil {
		return nil, errors.New("the database is nil") // Returns nil pointer. Henry's Way
	}
	repository := OrderRepository{}
	repository.database = gormDB // nil can be
	return &repository, nil
}

func (repository *OrderRepository) FetchAll(page int, ordersPerPage int) ([]models.Order, int, error) {

	var orders []models.Order

	var ordersCount int64
	if err := repository.database.Table("users").Count(&ordersCount).Error; err != nil {
		return nil, 0, err
	}

	numberOfPages := float64(ordersCount) / float64(ordersPerPage) // Calculates the number of pages of offers that we have.
	pageCount := int(math.Ceil(numberOfPages))                     // Rounds up the result of the numberOfOffers / offersPerPage
	if pageCount == 0 {
		pageCount = 1
	}

	offset := (page - 1) * ordersPerPage
	if err := repository.database.Limit(ordersPerPage).Offset(offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, pageCount, nil
}
func (repository *OrderRepository) Create(order models.Order) (models.Order, error) {
	order.ID = 0 // Remove the possibility of giving the ID in the request
	tx := repository.database.Create(&order)
	return order, tx.Error
}
