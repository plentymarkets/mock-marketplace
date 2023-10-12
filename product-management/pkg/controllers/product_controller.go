package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management/pkg/repositories"
)

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
		page := 1
		products, err, pageCount := controller.productRepository.GetProducts(page)

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
