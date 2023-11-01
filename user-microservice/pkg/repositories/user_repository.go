package repositories

import (
	"gorm.io/gorm"
	"log"
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

func (repository *UserRepository) GetUser(id int) models.User {
	var user models.User
	repository.Database.First(&user, id)
	return user
}

func (repository *UserRepository) UpdateUser(user models.User) {
	repository.Database.Save(&user)
}

func (repository *UserRepository) DeleteUser(user models.User) {
	repository.Database.Delete(&user)
}

func (repository *UserRepository) GetUserByEmail(email string) models.User {
	var user models.User
	tx := repository.Database.Where("email = ?", email).First(&user)

	if tx.Error != nil {
		log.Println(tx.Error)
		return models.User{}
	}

	return user
}
