package seed

import (
	"gorm.io/gorm"
	"user-microservice/pkg/models"
	"user-microservice/pkg/repositories"
)

func Seed(database *gorm.DB) error {
	userRepository := repositories.NewRepository(database)
	isEmpty := checkIfTableIsEmpty(userRepository)

	if isEmpty {
		user := generateUser()
		transaction := userRepository.Database.Create(&user)
		return transaction.Error
	}

	return nil
}

func generateUser() models.User {
	user := models.User{
		SellerID: 1,
		Email:    "john.doe@example.com",
		Password: "password",
	}

	return user
}

func checkIfTableIsEmpty(userRepository repositories.UserRepository) bool {
	var users []models.User
	userRepository.Database.Find(&users)
	return len(users) == 0
}
