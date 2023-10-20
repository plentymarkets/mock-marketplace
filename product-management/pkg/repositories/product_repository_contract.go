package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	FetchByID(id string) (models.Product, error)
	FetchAll(page int, productsPerPage int) ([]models.Product, int, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
}
