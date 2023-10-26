package routes

import (
	"auth/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
)

func RegisterRoutes(databaseConnection *gorm.DB) {
	router := gin.Default()

	api := router.Group("/auth")
	api.POST("/validate", controller.Auth(databaseConnection))

	err := router.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
