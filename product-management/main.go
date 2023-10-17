package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"product-management/pkg/controllers"
	"product-management/pkg/database"
	"product-management/pkg/middlewares"
	//"product-management/pkg/middlewares"
	"product-management/pkg/repositories"
)

func main() {

	engine := gin.Default()
	mariadb := database.NewMariaDBDatabase()

	// PRODUCT_CONTROLLER
	//------------------------------------------------------------------------------------------------------------------

	productRepository := repositories.NewProductRepository(mariadb.GetConnection())
	productController := controllers.NewProductController(&productRepository)

	product := engine.Group("/api/products").Use(middlewares.Authenticate())

	product.GET("/", productController.Get())
	product.GET("/:id", productController.GetByID())
	product.POST("/", productController.Create())
	product.PUT("/", productController.Update())
	product.DELETE("/:id", productController.Delete())
	//product.GET("/test", productController.GetProducts2)

	// VARIANT_CONTROLLER
	//------------------------------------------------------------------------------------------------------------------

	variantRepository := repositories.NewVariantRepository(mariadb.GetConnection())
	variantController := controllers.NewVariantController(&variantRepository)

	variant := engine.Group("/api/products").Use(middlewares.Authenticate())

	variant.GET("/", variantController.GetAll())
	variant.GET("/:id", variantController.GetByID())
	variant.POST("/", variantController.Create())
	variant.PUT("/", variantController.Update())
	variant.DELETE("/:id", variantController.Delete())

	err := engine.Run(":3004")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
