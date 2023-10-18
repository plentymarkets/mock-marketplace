package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	FetchByID(id string) (models.Product, error)
	FetchAll(page int) ([]models.Product, error, int)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id string)
}
