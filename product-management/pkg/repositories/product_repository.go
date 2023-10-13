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

func (repository *ProductRepository) GetProducts() ([]models.Product, error, string) {
	var products []models.Product
	repository.database.Find(&products)
	return products, nil, "pageCount"
}

func (repository *ProductRepository) CreateProduct(product models.Product) (models.Product, error) {
	repository.database.Create(&product)
	return product, nil
}

func (repository *ProductRepository) UpdateProduct(product models.Product) (models.Product, error) {
	repository.database.Model(&product).Updates(product)
	return product, nil
}

func (repository *ProductRepository) DeleteProduct(id string) {
	panic("Delete is not supported for this app")
}

func (repository *ProductRepository) GetProductByID(id string) (models.Product, error) {
	var product models.Product
	repository.database.Find(&product, id)
	return product, nil
}
