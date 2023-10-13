package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	GetProducts() ([]models.Product, error, string)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id string)
	GetProductByID(id string) (models.Product, error)
}
