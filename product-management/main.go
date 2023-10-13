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

	productRepository := repositories.NewProductRepository(mariadb.GetConnection())
	productController := controllers.NewProductController(&productRepository)

	product := engine.Group("/api/products").Use(middlewares.Authenticate())

	product.GET("/", productController.GetProducts())
	product.GET("/:id", productController.GetProductByID())
	product.POST("/", productController.CreateProduct())
	product.PUT("/", productController.UpdateProduct())
	product.DELETE("/:id", productController.DeleteProduct())

	err := engine.Run(":3004")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
