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

		c.JSON(http.StatusOK, map[string]any{
			"data":      products,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *ProductController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		gtin := c.Param("gtin")
		product, err := controller.productRepository.FetchByProduct(models.Product{GTIN: gtin})

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if product.ID == 0 {
			c.JSON(http.StatusOK, map[string]any{
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, map[string]any{
				"data": product,
			})
		}
		c.Done()
	}
}

func (controller *ProductController) Create() gin.HandlerFunc { // todo - If the VariantID is provided, it can change data in Variants !!!
	return func(c *gin.Context) {

		var product = models.Product{}
		err := c.BindJSON(&product)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		product, err = controller.productRepository.Create(product)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Success",
			"data":    product,
		})
		c.Done()
	}
}

func (controller *ProductController) Update() gin.HandlerFunc { // todo - investigate changes on the variant when changing the product.
	return func(c *gin.Context) {
		var product = models.Product{}
		err := c.BindJSON(&product)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		product, err = controller.productRepository.Update(product)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Product updated successfully",
			"data":    product,
		})
		c.Done()
	}
}

func (controller *ProductController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		gtin := c.Param("gtin")
		product, err := controller.productRepository.FetchByProduct(models.Product{GTIN: gtin})

		if product.ID == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Product not found!"})
			return
		}

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		product.Deleted = true

		product, err = controller.productRepository.Update(product)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		message, _ := fmt.Printf("Product with id %s Has been deleted successfully", gtin)

		c.JSON(http.StatusOK, map[string]any{
			"message": message,
			"data":    product,
		})
		c.Done()
	}
}
