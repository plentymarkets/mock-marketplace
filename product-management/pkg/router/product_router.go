package router

import (
	"github.com/gin-gonic/gin"
	"product-management/pkg/controllers"
	"product-management/pkg/database"
	"product-management/pkg/middlewares"
	"product-management/pkg/repositories"
)

func Product(mariadb database.MariaDBDatabase, engine *gin.Engine) {
	productRepository := repositories.NewProductRepository(mariadb.GetConnection())
	productController := controllers.NewProductController(&productRepository)

	product := engine.Group("/api/product").Use(middlewares.Authenticate())

	product.GET("/", productController.GetAll())
	product.GET("/:id", productController.GetByID())
	product.POST("/", productController.Create())
	product.PUT("/", productController.Update())
	product.DELETE("/:id", productController.Delete())
	//product.GET("/test", productController.GetProducts2)
}
