package repositories

import (
	"gorm.io/gorm"
	"product-management/pkg/models"
)

type VariantRepository struct {
	database *gorm.DB
}

func NewVariantRepository(gormDB *gorm.DB) VariantRepository {
	repository := VariantRepository{}
	repository.database = gormDB
	return repository
}

func (repository *VariantRepository) FetchAllByID(id string) (models.Variant, error) {
	var variant models.Variant
	repository.database.Model(&models.Variant{}).Preload("Variants").Find(&variant, id)
	return variant, nil
}

func (repository *VariantRepository) FetchAll() ([]models.Variant, error, string) {
	var variants []models.Variant
	repository.database.Model(&models.Variant{}).Preload("Variants").Find(&variants)
	return variants, nil, "pageCount"
}

func (repository *VariantRepository) Create(variant models.Variant) (models.Variant, error) {
	repository.database.Create(&variant)
	return variant, nil
}

func (repository *VariantRepository) Update(variant models.Variant) (models.Variant, error) {
	repository.database.Model(&variant).Updates(variant)
	return variant, nil
}

func (repository *VariantRepository) Delete(id string) {
	_ = id
	panic("Delete is not supported for this app")
}
