package repositories

import (
	"offer-management/pkg/models"
)

type OfferRepositoryContract interface {
	FetchByID(id string) (models.Offer, error)
	FetchAll(page int, offersPerPage int) ([]models.Offer, int, error)
	Create(offer models.Offer) (models.Offer, error)
	Update(offer models.Offer) (models.Offer, error)
}
