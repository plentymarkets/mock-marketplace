package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"offer-management/pkg/models"
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

func (repository *UserRepository) FetchAll(page int, usersPerPage int) ([]models.User, int, error) {

	var users []models.User

	var userCount int64
	if err := repository.database.Table("users").Count(&userCount).Error; err != nil {
		return nil, 0, err
	}

	numberOfPages := float64(userCount) / float64(usersPerPage) // Calculates the number of pages of offers that we have.
	pageCount := int(math.Ceil(numberOfPages))                  // Rounds up the result of the numberOfOffers / offersPerPage
	if pageCount == 0 {
		pageCount = 1
	}

	offset := (page - 1) * usersPerPage
	if err := repository.database.Limit(usersPerPage).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, pageCount, nil
}

func (repository *UserRepository) FetchByID(id string) (models.User, error) {
	var user models.User
	tx := repository.database.First(&user, id)
	return user, tx.Error
}

func (repository *UserRepository) FetchByName(UserName string) (models.User, error) {
	var user models.User
	tx := repository.database.Where("user_name = ?", UserName).First(&user)
	return user, tx.Error
}

func (repository *UserRepository) Create(user models.User) (models.User, error) {
	user.ID = 0 // Remove the possibility of giving the ID in the request
	tx := repository.database.Create(&user)
	return user, tx.Error
}

func (repository *UserRepository) Update(user models.User) (models.User, error) {
	tx := repository.database.Updates(&user)
	return user, tx.Error
}
