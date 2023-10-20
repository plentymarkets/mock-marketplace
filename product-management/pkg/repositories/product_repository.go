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

func NewProductRepository(gormDB *gorm.DB) (*ProductRepository, error) {
	// if the database is nil, it will crash. Throw an error
	if gormDB == nil {
		return nil, errors.New("the database is nil") // Returns nil pointer. Henry's Way
	}
	repository := ProductRepository{}
	repository.database = gormDB // nil can be
	return &repository, nil
}

func (repository *ProductRepository) FetchByID(id string) (models.Product, error) {
	var product models.Product
	tx := repository.database.Model(&models.Product{}).Preload("Variants").Find(&product, id)
	return product, tx.Error
}

func (repository *ProductRepository) FetchAll(page int, productsPerPage int) ([]models.Product, int, error) {

	var products []models.Product

	var productCount int64
	if err := repository.database.Table("products").Count(&productCount).Error; err != nil {
		return nil, 0, err
	}

	numberOfPages := float64(productCount) / float64(productsPerPage) // Calculates the number of pages of products that we have.
	pageCount := int(math.Ceil(numberOfPages))                        // Rounds up the result of the numberOfProducts / productsPerPage
	if pageCount == 0 {
		pageCount = 1
	}

	offset := (page - 1) * productsPerPage
	if err := repository.database.Limit(productsPerPage).Offset(offset).
		Preload("Variants").Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, pageCount, nil
}

func (repository *ProductRepository) Create(product models.Product) (models.Product, error) {
	repository.database.Create(&product) // Todo Error handling
	return product, nil
}

func (repository *ProductRepository) Update(product models.Product) (models.Product, error) {
	repository.database.Model(&product).Updates(product) // Todo Error handling
	return product, nil
}
