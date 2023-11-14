package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"os"
	"user-microservice/pkg/controller"
)

func RegisterRoutes(databaseConnection *gorm.DB) {
	router := gin.Default()

	api := router.Group("/user")
	api.GET("/validation", controller.Validate(databaseConnection))
	api.POST("/token", controller.RetrieveToken(databaseConnection))

	address := fmt.Sprintf("%s:%s",
		os.Getenv("GIN_HOST"),
		os.Getenv("GIN_PORT"),
	)
	err := router.Run(address)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
