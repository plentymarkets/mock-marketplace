package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"offer-management/pkg/controllers"
	"offer-management/pkg/middlewares"
	"offer-management/pkg/repositories"
)

func Offer(mariadb *gorm.DB, engine *gin.Engine) {

	// Homework = Check what the Engin
	offerRepository, _ := repositories.NewOfferRepository(mariadb)
	productRepository, _ := repositories.NewProductRepository(mariadb)
	offerController := controllers.NewOfferController(offerRepository, productRepository)

	offer := engine.Group("/offer").Use(middlewares.Authenticate())

	offer.GET("/", offerController.GetAll())
	offer.GET("/:id", offerController.GetByID())
	offer.POST("/", offerController.Create())
	offer.PUT("/:id", offerController.Update())
	offer.DELETE("/:id", offerController.Delete())
}
