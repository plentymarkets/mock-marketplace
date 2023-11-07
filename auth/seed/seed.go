package seed

import (
	"auth/models"
	"auth/repositories"
	"gorm.io/gorm"
)

func Seed(database *gorm.DB) {
	userRepository := repositories.NewRepository(database)
	user := generateUser()
	userRepository.CreateUser(user)
}

func generateUser() models.User {
	user := models.User{
		Email:    "john.doe@example.com",
		Password: "password",
	}

	return user
}
