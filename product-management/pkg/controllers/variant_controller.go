package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management/pkg/models"
	"product-management/pkg/repositories"
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
		variants, err, pageCount := controller.variantRepository.FetchAll()

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"data":      variants,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *VariantController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		variant, err := controller.variantRepository.FetchAllByID(id)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"data": variant,
		})
		c.Done()
	}
}

func (controller *VariantController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var variant = models.Variant{}
		err := c.BindJSON(&variant)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		variant, err = controller.variantRepository.Create(variant)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Success",
			"data":    variant,
		})
		c.Done()
	}
}

func (controller *VariantController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var variant = models.Variant{}
		err := c.BindJSON(&variant)

		variant, err = controller.variantRepository.Update(variant)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Variant updated successfully",
			"data":    variant,
		})
		c.Done()
	}
}

func (controller *VariantController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		variant, _ := controller.variantRepository.FetchAllByID(id)

		variant, err := controller.variantRepository.Update(variant)
		if err != nil {

			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"pageCount": "Variant with id " + id + " Has been deleted successfully",
			"data":      variant,
		})
		c.Done()
	}
}
