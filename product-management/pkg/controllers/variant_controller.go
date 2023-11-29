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

type VariantController struct {
	variantRepository repositories.VariantRepositoryContract
}

func NewVariantController(variantRepository repositories.VariantRepositoryContract) VariantController {
	return VariantController{
		variantRepository: variantRepository,
	}
}

func (controller *VariantController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)

		if err != nil {
			log.Printf("Invalid page number unsupported format %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid page number format! Page number should be an integer."})
			return
		}

		variants, pageCount, err := controller.variantRepository.FetchAll(page, 10)

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
			"data":      variants,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *VariantController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.Atoi(id)

		if err != nil {
			log.Printf("Invalid product ID %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		variant, err := controller.variantRepository.FetchById(id)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, variant)
		c.Done()
	}
}

func (controller *VariantController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var variant = models.Variant{}
		err := c.BindJSON(&variant)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		variant, err = controller.variantRepository.Create(variant)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, variant)
		c.Done()
	}
}

func (controller *VariantController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatedVariant = models.Variant{}
		err := c.BindJSON(&updatedVariant)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		id := c.Param("id")
		existingVariant, err := controller.variantRepository.FetchById(id)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if existingVariant.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid request"})
			return
		}

		_, err = controller.variantRepository.Update(existingVariant, updatedVariant)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Variant updated successfully"})
		c.Done()
	}
}

func (controller *VariantController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		variant, err := controller.variantRepository.FetchById(id)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if variant.ID == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Invalid request"})
			return
		}

		if variant.Deleted == true {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Variant is already Deleted"})
			return
		}

		// Set variant as deleted
		variant.Deleted = true

		variant, err = controller.variantRepository.Update(variant, variant)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Variant with id " + id + " Has been deleted successfully"})
		c.Done()
	}
}
