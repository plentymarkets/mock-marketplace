package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"product-management/pkg/models"
	"product-management/pkg/repositories"
	"strconv"
)

const ProductsPerPage = 10

type ProductController struct {
	productRepository repositories.ProductRepositoryContract
}

func NewProductController(productRepository repositories.ProductRepositoryContract) ProductController {
	return ProductController{
		productRepository: productRepository,
	}
}

func (controller *ProductController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)

		if err != nil {
			log.Printf("Invalid page number unsupported format %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid page number format! Page number should be an integer."})
			return
		}

		// TODO - Implement this
		//username := c.DefaultQuery("user", "")
		//if username == "" {
		//	uuid := mdHashing(username)
		//	products, pageCount, err := controller.productRepository.GetProductsByToken(uuid)
		//}

		products, pageCount, err := controller.productRepository.FetchAll(page, ProductsPerPage)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if page < 1 || page > pageCount {
			log.Println("Invalid page number!")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid page number! Please selet a page from 1 to %d", pageCount)})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":      products,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *ProductController) GetByGTIN() gin.HandlerFunc {
	return func(c *gin.Context) {
		gtin := c.Param("gtin")
		product, err := controller.productRepository.FetchProductByGTIN(gtin)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if product.ID == 0 {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, product)
		}
		c.Done()
	}
}

func (controller *ProductController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var product = models.Product{}
		err := c.BindJSON(&product)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		product, err = controller.productRepository.Create(product, c.Request.Header.Get("token"))

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, product)
		c.Done()
	}
}

func (controller *ProductController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updatedProduct = models.Product{}
		err := c.BindJSON(&updatedProduct)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		gtin := c.Param("gtin")
		existingProduct, err := controller.productRepository.FetchProductByGTIN(gtin)

		if existingProduct.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found!"})
			return
		}

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		_, err = controller.productRepository.Update(existingProduct, updatedProduct)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
		c.Done()
	}
}

func (controller *ProductController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		gtin := c.Param("gtin")
		token := c.Request.Header.Get("token")

		product, err := controller.productRepository.GetProductByTokenAndGTIN(token, gtin)

		if product.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found!"})
			return
		}

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		product.Deleted = true

		product, err = controller.productRepository.Update(product, product)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Product with GTIN: %s has been deleted successfully", gtin)})
		c.Done()
	}
}
