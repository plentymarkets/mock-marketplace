package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
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

func (repository *ProductRepository) FetchAll(page int) ([]models.Product, error, int) {

	var products []models.Product

	var ordersCount int64
	if err := repository.database.Table("products").Count(&ordersCount).Error; err != nil {
		return nil, err, 0
	}

	const productsPerPage = 3
	pageCount := int(math.Ceil(float64(ordersCount) / float64(productsPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}

	if page < 1 || page > pageCount {
		return nil, errors.New(fmt.Sprintf("invalid page number the page should be grater than 0 and lower than %d", pageCount+1)), 0
	}

	offset := (page - 1) * productsPerPage
	if err := repository.database.Limit(productsPerPage).Offset(offset).
		Preload("Variants").Find(&products).Error; err != nil {
		return nil, err, 0
	}

	return products, nil, pageCount
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
