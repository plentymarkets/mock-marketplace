package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	GetProducts(page int) ([]models.Product, error, int)
	CreateProduct(product models.Product)
}
