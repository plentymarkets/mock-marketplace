package repositories

import (
	"errors"
	"gorm.io/gorm"
	"offer-microservice/pkg/models"
)

type OfferRepository struct {
	Database *gorm.DB
}

func NewOfferRepository(databaseConnection *gorm.DB) OfferRepository {
	repository := OfferRepository{}
	repository.Database = databaseConnection
	return repository
}

//func (offerRepository OfferRepository) FindAll(offset *int, limit *int) (*[]models.Offer, error) {
//
//	if offset != nil {
//		offerRepository.Database = offerRepository.Database.Offset(*offset)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	if limit != nil {
//		offerRepository.Database = offerRepository.Database.Limit(*limit)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	var offers []models.Offer
//	offerRepository.Database.Find(&offers)
//
//	if offerRepository.Database.Error != nil {
//		return nil, offerRepository.Database.Error
//	}
//
//	return &offers, nil
//}

//func (offerRepository OfferRepository) FindById(id int, offset *int, limit *int) (*models.Offer, error) {
//	var offer models.Offer
//
//	if offset != nil {
//		offerRepository.Database = offerRepository.Database.Offset(*offset)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	if limit != nil {
//		offerRepository.Database = offerRepository.Database.Limit(*limit)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	offerRepository.Database.Find(&offer, id)
//
//	return &offer, nil
//}

//func (offerRepository OfferRepository) FindOneById(id int) (*models.Offer, error) {
//
//	var offer models.Offer
//	offerRepository.Database.First(&offer, id)
//
//	if offerRepository.Database.Error != nil {
//		return nil, offerRepository.Database.Error
//	}
//
//	return &offer, nil
//}

func (offerRepository OfferRepository) FindOneByField(field string, value string) (*models.Offer, error) {
	var offer models.Offer

	offerRepository.Database.Where(field+" = ?", value).First(&offer)

	if offerRepository.Database.Error != nil && !errors.Is(offerRepository.Database.Error, gorm.ErrRecordNotFound) {
		return nil, offerRepository.Database.Error
	}

	if offerRepository.Database.RowsAffected == 0 {
		return nil, nil
	}

	return &offer, nil
}

func (offerRepository OfferRepository) FindOneByFields(fields map[string]string) (*models.Offer, error) {
	var offer models.Offer

	for key, val := range fields {
		offerRepository.Database = offerRepository.Database.Where(key+" = ?", val)

		if offerRepository.Database.Error != nil {
			return nil, offerRepository.Database.Error
		}
	}

	offerRepository.Database.First(&offer)

	if offerRepository.Database.Error != nil && !errors.Is(offerRepository.Database.Error, gorm.ErrRecordNotFound) {
		return nil, offerRepository.Database.Error
	}

	if offerRepository.Database.RowsAffected == 0 {
		return nil, nil
	}

	return &offer, nil
}

func (offerRepository OfferRepository) FindByField(field string, value string, offset *int, limit *int) (*[]models.Offer, error) {
	//var offers []models.Offer
	offers := make([]models.Offer, 0, *limit)

	if offset != nil {
		offerRepository.Database = offerRepository.Database.Offset(*offset)

		if offerRepository.Database.Error != nil {
			return nil, offerRepository.Database.Error
		}
	}

	if limit != nil {
		offerRepository.Database = offerRepository.Database.Limit(*limit)

		if offerRepository.Database.Error != nil {
			return nil, offerRepository.Database.Error
		}
	}

	offerRepository.Database.Where(field+" = ?", value).Find(&offers)

	if offerRepository.Database.Error != nil && !errors.Is(offerRepository.Database.Error, gorm.ErrRecordNotFound) {
		return nil, offerRepository.Database.Error
	}

	if offerRepository.Database.RowsAffected == 0 {
		return nil, nil
	}

	return &offers, nil
}

//func (offerRepository OfferRepository) FindByFields(fields map[string]string, offset *int, limit *int) (*[]models.Offer, error) {
//	var offers []models.Offer
//
//	if offset != nil {
//		offerRepository.Database = offerRepository.Database.Offset(*offset)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	if limit != nil {
//		offerRepository.Database = offerRepository.Database.Limit(*limit)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	for key, val := range fields {
//		offerRepository.Database = offerRepository.Database.Where(key+" = ?", val)
//
//		if offerRepository.Database.Error != nil {
//			return nil, offerRepository.Database.Error
//		}
//	}
//
//	offerRepository.Database.Find(&offers)
//
//	if offerRepository.Database.Error != nil && !errors.Is(offerRepository.Database.Error, gorm.ErrRecordNotFound) {
//		return nil, offerRepository.Database.Error
//	}
//
//	if offerRepository.Database.RowsAffected == 0 {
//		return nil, nil
//	}
//
//	return &offers, nil
//}

func (offerRepository OfferRepository) Create(offer models.Offer) (models.Offer, error) {
	transaction := offerRepository.Database.Create(&offer)
	return offer, transaction.Error
}
