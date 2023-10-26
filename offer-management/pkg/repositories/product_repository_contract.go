package repositories

import (
	"offer-management/pkg/models"
)

type ProductRepositoryContract interface {
	FetchByProduct(models.Product) (models.Product, error)
	FetchAll(page int, productsPerPage int) ([]models.Product, int, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
}
