package seed

import (
	"gorm.io/gorm"
	"offer-microservice/pkg/models"
	"offer-microservice/pkg/repositories"
)

func Seed(databaseConnection *gorm.DB) error {
	offerRepository := repositories.NewOfferRepository(databaseConnection)
	isEmpty := checkIfTableIsEmpty(offerRepository)

	if isEmpty {
		offer := generateOffer()
		transaction := offerRepository.Database.Create(&offer)
		return transaction.Error
	}

	return nil
}

func generateOffer() models.Offer {
	offer := models.Offer{
		SellerID:  1,
		ProductID: 1,
		Price:     "9.99",
		Quantity:  12,
	}

	return offer
}

func checkIfTableIsEmpty(offerRepository repositories.OfferRepository) bool {
	var offers []models.Offer
	offerRepository.Database.Find(&offers)
	return len(offers) == 0
}
