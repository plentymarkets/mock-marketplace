package migrate

import (
	"auth/helper"
	"auth/models"
	"fmt"
)

func Migrate() {
	databaseConnection := helper.GetDatabaseConnection()

	fmt.Println(&models.User{})

	err := databaseConnection.AutoMigrate(&models.User{})

	if err != nil {
		panic("Could not migrate database")
	}
}
