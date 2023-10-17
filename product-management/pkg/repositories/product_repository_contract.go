package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	FetchAll() ([]models.Product, error, string)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id string)
	FetchByID(id string) (models.Product, error)
}
