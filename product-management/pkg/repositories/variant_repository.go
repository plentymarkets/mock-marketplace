package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"product-management/pkg/models"
)

type VariantRepository struct {
	database *gorm.DB
}

func NewVariantRepository(gormDB *gorm.DB) (VariantRepository, error) {
	if gormDB == nil {
		return VariantRepository{}, errors.New("the database is nil") // Returns empty struct. Max's Way
	}
	repository := VariantRepository{}
	repository.database = gormDB
	return repository, nil
}

func (repository *VariantRepository) FetchAll(page int, variantsPerPage int) ([]models.Variant, int, error) {
	var variants []models.Variant

	var variantCount int64
	if err := repository.database.Table("variants").Count(&variantCount).Error; err != nil {
		return nil, 0, err
	}

	// Calculates the number of pages of variants that we have.
	numberOfPages := float64(variantCount) / float64(variantsPerPage)
	// Rounds up the result of the numberOfProducts / variantsPerPage
	pageCount := int(math.Ceil(numberOfPages))
	if pageCount == 0 {
		pageCount = 1
	}

	offset := (page - 1) * variantsPerPage
	if err := repository.database.Limit(variantsPerPage).Offset(offset).
		Preload("Variants").Find(&variants).Error; err != nil {
		return nil, 0, err
	}
	return variants, pageCount, nil
}

func (repository *VariantRepository) FetchById(id string) (models.Variant, error) {
	var variants models.Variant
	repository.database.Model(&models.Variant{}).Find(&variants, id)

	return variants, nil
}

func (repository *VariantRepository) FetchByProductId() ([]models.Variant, error) {
	var variants []models.Variant
	repository.database.Model(&models.Variant{}).Preload("Variants").Find(&variants)
	return variants, nil
}

func (repository *VariantRepository) Create(variant models.Variant) (models.Variant, error) {
	repository.database.Create(&variant)
	return variant, nil
}

func (repository *VariantRepository) Update(variant models.Variant) (models.Variant, error) {
	repository.database.Model(&variant).Updates(variant)
	return variant, nil
}
