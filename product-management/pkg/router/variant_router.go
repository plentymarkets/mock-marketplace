package router

import (
	"github.com/gin-gonic/gin"
	"product-management/pkg/controllers"
	"product-management/pkg/database"
	"product-management/pkg/middlewares"
	"product-management/pkg/repositories"
)

func Variant(mariadb database.MariaDBDatabase, engine *gin.Engine) {

	variantRepository := repositories.NewVariantRepository(mariadb.GetConnection())
	variantController := controllers.NewVariantController(&variantRepository)

	variant := engine.Group("/api/variant").Use(middlewares.Authenticate())

	variant.GET("/", variantController.GetAll())
	variant.GET("/:id", variantController.GetByID())
	variant.POST("/", variantController.Create())
	variant.PUT("/", variantController.Update())
	variant.DELETE("/:id", variantController.Delete())
}
