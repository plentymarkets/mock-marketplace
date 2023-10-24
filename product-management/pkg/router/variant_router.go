package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"product-management/pkg/controllers"
	"product-management/pkg/middlewares"
	"product-management/pkg/repositories"
)

func Variant(mariadb *gorm.DB, engine *gin.Engine) {

	variantRepository, _ := repositories.NewVariantRepository(mariadb)
	variantController := controllers.NewVariantController(&variantRepository)

	variant := engine.Group("/api/variant").Use(middlewares.Authenticate())

	variant.GET("/", variantController.GetAll())
	variant.GET("/:id", variantController.GetByID())
	variant.POST("/", variantController.Create())
	variant.PUT("/", variantController.Update())
	variant.DELETE("/:id", variantController.Delete())
}
