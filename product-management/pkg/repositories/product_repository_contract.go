package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	FetchByID(id string) (models.Product, error)
	FetchAll() ([]models.Product, error, string)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id string)
}
