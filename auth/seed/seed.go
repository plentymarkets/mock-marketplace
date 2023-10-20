package seed

import (
	"auth/models"
	"auth/repositories"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
}

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
