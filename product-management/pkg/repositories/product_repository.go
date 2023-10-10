package repositories

import (
	"gorm.io/gorm"
	"product-management/pkg/models"
)

type ProductRepository struct {
	database *gorm.DB
}

func NewProductRepository(gormDB *gorm.DB) ProductRepository {
	repository := ProductRepository{}
	repository.database = gormDB
	return repository
}

func (repository *ProductRepository) GetProducts(page int) ([]models.Product, error, int) {
	// get all products.
	return nil, nil, 0
}

func (repository *ProductRepository) CreateProduct(product models.Product) {
	repository.database.Create(&product)
}
