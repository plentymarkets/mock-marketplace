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

func (repository *ProductRepository) FetchByID(id string) (models.Product, error) {
	var product models.Product
	repository.database.Model(&models.Product{}).Preload("Variants").Find(&product, id)
	return product, nil
}

func (repository *ProductRepository) FetchAll() ([]models.Product, error, string) {
	var products []models.Product
	repository.database.Model(&models.Product{}).Preload("Variants").Find(&products)
	return products, nil, "pageCount"
}

func (repository *ProductRepository) Create(product models.Product) (models.Product, error) {
	repository.database.Create(&product)
	return product, nil
}

func (repository *ProductRepository) Update(product models.Product) (models.Product, error) {
	repository.database.Model(&product).Updates(product)
	return product, nil
}

func (repository *ProductRepository) Delete(id string) {
	_ = id
	panic("Delete is not supported for this app")
}
