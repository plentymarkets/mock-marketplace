package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"product-management/pkg/controllers"
	"product-management/pkg/middlewares"
	"product-management/pkg/repositories"
)

func Product(mariadb *gorm.DB, engine *gin.Engine) {

	// Homework = Check what the Engin
	productRepository, _ := repositories.NewProductRepository(mariadb)
	productController := controllers.NewProductController(productRepository)

	product := engine.Group("/api/product").Use(middlewares.Authenticate())

	product.GET("/", productController.GetAll())
	product.GET("/:id", productController.GetByID())
	product.POST("/", productController.Create())
	product.PUT("/", productController.Update())
	product.DELETE("/:id", productController.Delete())
	//product.GET("/test", productController.GetProducts2)
}
