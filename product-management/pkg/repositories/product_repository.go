package repositories

import (
	"errors"
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

func (repository *ProductRepository) GetProducts(page int) ([]models.Product, error, int) {
	var products []models.Product

	var productCount int64
	if err := repository.database.Table("products").Count(&productCount).Error; err != nil {
		return nil, err, 0
	}

	const productsPerPage = 15
	pageCount := int(math.Ceil(float64(productCount) / float64(productsPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}

	if page < 1 || page > pageCount {
		return nil, errors.New("bad request"), 0
	}

	offset := (page - 1) * productsPerPage
	if err := repository.database.Limit(productsPerPage).Offset(offset).
		Preload("Variants").Find(&products).Error; err != nil {
		return nil, err, 0
	}

	return products, nil, pageCount
}

func (repository *ProductRepository) CreateProduct(product models.Product) {
	repository.database.Create(&product)
}
