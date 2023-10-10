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

	r := gin.Default()

	mariadb := database.NewMariaDBDatabase()

	productRepository := repositories.NewProductRepository(mariadb.GetConnection())
	productController := controllers.NewProductController(&productRepository)

	product := r.Group("/api/products").Use(middlewares.Authenticate())

	product.GET("/", productController.GetProducts())

	//products.GET("/:id", productController.Getproducts())
	//products.POST("/", productController.Getproducts())
	//products.PUT("/", productController.Getproducts())
	//products.DELETE("/", productController.Getproducts())

	err := r.Run(":3004")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func main2() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, World!")
	//})
}
