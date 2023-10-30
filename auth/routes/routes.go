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

	api := router.Group("/user")
	api.POST("/validation", controller.Validate())
	api.POST("/token", controller.RetrieveToken(databaseConnection))

	err := router.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
