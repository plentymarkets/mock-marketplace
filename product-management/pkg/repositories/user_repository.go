package repositories

import (
	"errors"
	"gorm.io/gorm"
	"product-management/pkg/models"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) (*UserRepository, error) {
	// if the database is nil, it will crash. Throw an error
	if gormDB == nil {
		return nil, errors.New("the database is nil") // Returns nil pointer. Henry's Way
	}
	repository := UserRepository{}
	repository.database = gormDB // nil can be
	return &repository, nil
}

func (repository *UserRepository) FetchByUser(models.User) (models.User, error) {
	var user = models.User{UUID: "098f6bcd4621d373cade4e832627b4f16"}
	tx := repository.database.Where(&user).Find(&user)
	return user, tx.Error
}

func (repository *UserRepository) Create(user models.User) (models.User, error) {
	user.ID = 0 // Remove the possibility of giving the ID in the request
	tx := repository.database.Create(&user)
	return user, tx.Error
}

func (repository *UserRepository) Update(user models.User) (models.User, error) {
	tx := repository.database.Model(&user).Updates(user)
	return user, tx.Error
}
