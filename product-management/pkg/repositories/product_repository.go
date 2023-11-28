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
	tx := repository.database.Preload("Variants").Find(&product, id)
	return product, tx.Error
}

func (repository *ProductRepository) FetchByProduct(product models.Product) (models.Product, error) {
	var retrievedProduct models.Product
	tx := repository.database.Where(product).Preload("Variants").Find(&retrievedProduct)
	return retrievedProduct, tx.Error
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

func (repository *ProductRepository) Create(product models.Product, userToken string) (models.Product, error) {

	// TODO - Move this to User.
	var user models.User
	err := repository.database.Where("token = ?", userToken).First(&user).Error

	if user.ID == 0 || err != nil {
		return product, err
	}

	product.ID = 0 // Remove the possibility of giving the ID in the request
	product.UserID = user.ID

	err = repository.database.Create(&product).Error
	return product, err
}

func (repository *ProductRepository) Update(existingProduct models.Product, updatedProduct models.Product) (models.Product, error) {
	tx := repository.database.Model(&existingProduct).Updates(updatedProduct)
	return existingProduct, tx.Error
}

func (repository *ProductRepository) GetProductByTokenAndGTIN(token string, gtin string) (models.Product, error) {
	var products models.Product

	err := repository.database.
		Joins("JOIN users u ON products.user_id = u.id").
		Joins("JOIN variants v ON v.product_id = products.id").
		Where("v.gtin = ? AND u.token = ?", gtin, token).
		Preload("Variants").
		FirstOrInit(&products).Error

	return products, err
}

func (repository *ProductRepository) FetchProductByGTIN(gtin string) (models.Product, error) {
	var product models.Product

	err := repository.database.
		Joins("JOIN users u ON products.user_id = u.id").
		Joins("JOIN variants v ON v.product_id = products.id").
		Where("v.gtin = ?", gtin).
		Preload("Variants", "gtin = ?", gtin). // Filter variants based on GTIN
		FirstOrInit(&product)

	return product, err.Error
}
