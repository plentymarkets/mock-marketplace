package repositories

import (
	"auth/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewRepository(databaseConnection *gorm.DB) UserRepository {
	repository := UserRepository{}
	repository.database = databaseConnection
	return repository
}

func (repository *UserRepository) CreateUser(order models.User) {
	repository.database.Create(&order)
}

func (repository *UserRepository) GetUser(id int) models.User {
	var user models.User
	repository.database.First(&user, id)
	return user
}

func (repository *UserRepository) UpdateUser(user models.User) {
	repository.database.Save(&user)
}

func (repository *UserRepository) DeleteUser(user models.User) {
	repository.database.Delete(&user)
}

func (repository *UserRepository) GetUserByEmail(email string) models.User {
	var user models.User
	repository.database.Where("email = ?", email).First(&user) // Error handling

	return user
}
