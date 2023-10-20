package repositories

import (
	"product-management/pkg/models"
)

type VariantRepositoryContract interface {
	FetchAll(page int, variantsPerPage int) ([]models.Variant, int, error)
	FetchById(id string) (models.Variant, error)
	FetchByProductId() ([]models.Variant, error)
	Create(variant models.Variant) (models.Variant, error)
	Update(variant models.Variant) (models.Variant, error)
}
