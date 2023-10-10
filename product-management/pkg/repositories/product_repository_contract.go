package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	GetProducts(page int) ([]models.Product, error, int)
	CreateProducts(product models.Product)
}
