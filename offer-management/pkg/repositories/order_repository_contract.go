package repositories

import (
	"offer-management/pkg/models"
)

type OrderRepositoryContract interface {
	FetchAll(page int, OrdersPerPage int) ([]models.Order, int, error)
	Create(order models.Order) (models.Order, error)
}
