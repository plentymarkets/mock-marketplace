package repositories

import (
	"product-management/pkg/models"
)

type ProductRepositoryContract interface {
	FetchByID(id string) (models.Product, error)
	FetchByProduct(product models.Product) (models.Product, error)
	FetchAll(page int, productsPerPage int) ([]models.Product, int, error)
	Create(product models.Product, token string) (models.Product, error)
	Update(existingProduct models.Product, updatedProduct models.Product) (models.Product, error)
	GetProductByTokenAndGTIN(token string, gtin string) (models.Product, error)
	FetchProductByGTIN(gtin string) (models.Product, error)
}
