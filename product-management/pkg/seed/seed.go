package seed

import (
	"gorm.io/gorm"
	"product-management/pkg/models"
	"product-management/pkg/repositories"
)

func Seed(databaseConnection *gorm.DB) error {
	productRepository, err := repositories.NewProductRepository(databaseConnection)

	if err != nil {
		return err
	}

	isEmpty := checkIfTableIsEmpty(*productRepository)

	if isEmpty {
		product := generateProduct()
		_, err := productRepository.Create(product)
		return err
	}

	return nil
}

func generateProduct() models.Product {
	product := models.Product{
		UserID:        1,
		Name:          "Product Name",
		Description:   "Product Description",
		GTIN:          "123456789",
		Categories:    1,
		Manufacturers: 1,
		Attributes:    "Red, Large, Leather",
		Deleted:       false,
	}

	for i := 0; i < 3; i++ {
		variant := generateVariant(product)
		product.Variants = append(product.Variants, variant)
	}

	return product
}

func generateVariant(product models.Product) models.Variant {
	orderItem := models.Variant{
		ProductID:  product.ID,
		UserID:     1,
		Name:       "Variant Name",
		GTIN:       "987654321",
		Attributes: "Blue, Small, Plastic",
		Deleted:    false,
	}

	return orderItem
}

func checkIfTableIsEmpty(productRepository repositories.ProductRepository) bool {
	var products []models.Product
	var err error
	products, _, err = productRepository.FetchAll(1, 1)

	if err != nil {
		return false
	}

	return len(products) == 0
}
