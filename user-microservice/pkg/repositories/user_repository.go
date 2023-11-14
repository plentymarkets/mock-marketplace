package repositories

import (
	"gorm.io/gorm"
	"user-microservice/pkg/models"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewRepository(databaseConnection *gorm.DB) UserRepository {
	repository := UserRepository{}
	repository.Database = databaseConnection
	return repository
}

func (repository *UserRepository) FindAll(offset *int, limit *int) (*[]models.User, error) {

	if offset != nil {
		repository.Database = repository.Database.Offset(*offset)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	if limit != nil {
		repository.Database = repository.Database.Limit(*limit)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	var offers []models.User
	repository.Database.Find(&offers)

	if repository.Database.Error != nil {
		return nil, repository.Database.Error
	}

	return &offers, nil
}

func (repository *UserRepository) FindById(id int, offset *int, limit *int) (*models.User, error) {
	var offer models.User

	if offset != nil {
		repository.Database = repository.Database.Offset(*offset)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	if limit != nil {
		repository.Database = repository.Database.Limit(*limit)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	repository.Database.Find(&offer, id)

	return &offer, nil
}

func (repository *UserRepository) FindOneById(id int) (*models.User, error) {

	var offer models.User
	repository.Database.First(&offer, id)

	if repository.Database.Error != nil {
		return nil, repository.Database.Error
	}

	return &offer, nil
}

func (repository *UserRepository) FindOneByField(field string, value string) (*models.User, error) {
	var offer models.User

	repository.Database.Where(field+" = ?", value).First(&offer)

	if repository.Database.Error != nil {
		return nil, repository.Database.Error
	}

	return &offer, nil
}

func (repository *UserRepository) FindOneByFields(fields map[string]string) (*models.User, error) {
	var offer models.User

	for key, val := range fields {
		repository.Database = repository.Database.Where(key+" = ?", val)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	repository.Database.First(&offer)

	if repository.Database.Error != nil {
		return nil, repository.Database.Error
	}

	return &offer, nil
}

func (repository *UserRepository) FindByField(field string, value string, offset *int, limit *int) (*[]models.User, error) {
	var offers []models.User

	if offset != nil {
		repository.Database = repository.Database.Offset(*offset)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	if limit != nil {
		repository.Database = repository.Database.Limit(*limit)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	repository.Database.Where(field+" = ?", value).Find(&offers)

	if repository.Database.Error != nil {
		return nil, repository.Database.Error
	}

	return &offers, nil
}

func (repository *UserRepository) FindByFields(fields map[string]string, offset *int, limit *int) (*[]models.User, error) {
	var offers []models.User

	if offset != nil {
		repository.Database = repository.Database.Offset(*offset)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	if limit != nil {
		repository.Database = repository.Database.Limit(*limit)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	for key, val := range fields {
		repository.Database = repository.Database.Where(key+" = ?", val)

		if repository.Database.Error != nil {
			return nil, repository.Database.Error
		}
	}

	repository.Database.Find(&offers)

	if repository.Database.Error != nil {
		return nil, repository.Database.Error
	}

	return &offers, nil
}

func (repository *UserRepository) Create(user *models.User) (*models.User, error) {
	transaction := repository.Database.Create(&user)
	return user, transaction.Error
}

func (repository *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	transaction := repository.Database.Save(&user)
	return user, transaction.Error
}

func (repository *UserRepository) DeleteUser(user *models.User) (*models.User, error) {
	transaction := repository.Database.Delete(&user)
	return user, transaction.Error
}
