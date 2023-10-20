package routes

import (
	"auth/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func RegisterRoutes(databaseConnection *gorm.DB) {
	router := gin.Default()
	api := router.Group("/api")
	api.POST("/auth", controller.Auth(databaseConnection))
	err := router.Run()

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
