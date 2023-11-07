package routes

import (
	"auth/controller"
	"fmt"
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

	address := fmt.Sprintf("%s:%s",
		os.Getenv("HOST"),
		os.Getenv("AUTH_PORT"),
	)
	err := router.Run(address)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
