package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"offer-management/pkg/models"
)

type OfferRepository struct {
	database *gorm.DB
}

func NewOfferRepository(gormDB *gorm.DB) (*OfferRepository, error) {
	// if the database is nil, it will crash. Throw an error
	if gormDB == nil {
		return nil, errors.New("the database is nil") // Returns nil pointer. Henry's Way
	}
	repository := OfferRepository{}
	repository.database = gormDB // nil can be
	return &repository, nil
}

func (repository *OfferRepository) FetchByID(id string) (models.Offer, error) {
	var offer models.Offer
	tx := repository.database.Model(&models.Offer{}).Find(&offer, id)
	return offer, tx.Error
}

func (repository *OfferRepository) FetchAll(page int, offersPerPage int) ([]models.Offer, int, error) {

	var offers []models.Offer

	var offerCount int64
	if err := repository.database.Table("offers").Count(&offerCount).Error; err != nil {
		return nil, 0, err
	}

	numberOfPages := float64(offerCount) / float64(offersPerPage) // Calculates the number of pages of offers that we have.
	pageCount := int(math.Ceil(numberOfPages))                    // Rounds up the result of the numberOfOffers / offersPerPage
	if pageCount == 0 {
		pageCount = 1
	}

	offset := (page - 1) * offersPerPage
	if err := repository.database.Limit(offersPerPage).Offset(offset).Find(&offers).Error; err != nil {
		return nil, 0, err
	}

	return offers, pageCount, nil
}

func (repository *OfferRepository) Create(offer models.Offer) (models.Offer, error) {
	offer.ID = 0 // Remove the possibility of giving the ID in the request
	tx := repository.database.Create(&offer)
	return offer, tx.Error
}

func (repository *OfferRepository) Update(offer models.Offer) (models.Offer, error) {
	tx := repository.database.Model(&offer).Updates(offer)
	return offer, tx.Error
}
