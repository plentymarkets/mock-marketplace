package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"offer-management/pkg/models"
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

func (repository *ProductRepository) FetchByProduct(product models.Product) (models.Product, error) {
	tx := repository.database.Where(&product).Preload("Offers").Find(&product)
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
	if err := repository.database.Limit(productsPerPage).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, pageCount, nil
}

func (repository *ProductRepository) Create(product models.Product) (models.Product, error) {
	product.ID = 0 // Remove the possibility of giving the ID in the request
	tx := repository.database.Create(&product)
	return product, tx.Error
}

func (repository *ProductRepository) Update(product models.Product) (models.Product, error) {
	tx := repository.database.Updates(&product)
	return product, tx.Error
}
