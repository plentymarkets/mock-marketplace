package repositories

import (
	"product-management/pkg/models"
)

type VariantRepositoryContract interface {
	FetchAllByID(id string) (models.Variant, error)
	FetchAll() ([]models.Variant, error, string)
	Create(variant models.Variant) (models.Variant, error)
	Update(variant models.Variant) (models.Variant, error)
	Delete(id string)
}
