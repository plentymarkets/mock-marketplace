package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management/pkg/models"
	"product-management/pkg/repositories"
	"time"
)

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ProductController struct {
	productRepository repositories.ProductRepositoryContract
}

func NewProductController(productRepository repositories.ProductRepositoryContract) ProductController {
	return ProductController{
		productRepository: productRepository,
	}
}

func (controller *ProductController) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err, pageCount := controller.productRepository.FetchAll()

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"data":      products,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *ProductController) GetProductByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product, err := controller.productRepository.FetchByID(id)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"data": product,
		})
		c.Done()
	}
}

func (controller *ProductController) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var product = models.Product{}
		err := c.BindJSON(&product)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		product, err = controller.productRepository.Create(product)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Success",
			"data":    product,
		})
		c.Done()
	}
}

func (controller *ProductController) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product = models.Product{}
		err := c.BindJSON(&product)

		product, err = controller.productRepository.Update(product)
		if err != nil {
			return
		}

		time.Sleep(100)

		c.JSON(http.StatusOK, map[string]any{
			"message": "Product updated successfully",
			"data":    product,
		})
		c.Done()
	}
}

func (controller *ProductController) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		product, _ := controller.productRepository.FetchByID(id)
		product.Deleted = true

		time.Sleep(100)
		product, err := controller.productRepository.Update(product)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"pageCount": "Product with id " + id + " Has been deleted successfully",
			"data":      product,
		})
		c.Done()
	}
}
