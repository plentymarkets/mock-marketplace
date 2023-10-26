package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"offer-management/pkg/models"
	"offer-management/pkg/repositories"
	"strconv"
	"time"
)

const OffersPerPage = 10

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type OfferController struct {
	offerRepository   repositories.OfferRepositoryContract
	productRepository repositories.ProductRepositoryContract
}

func NewOfferController(offerRepository repositories.OfferRepositoryContract, productRepository repositories.ProductRepositoryContract) OfferController {
	return OfferController{
		offerRepository:   offerRepository,
		productRepository: productRepository,
	}
}

func (controller *OfferController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		sku := c.DefaultQuery("sku", "")
		if sku != "" {
			var product = models.Product{SKU: sku}

			product, err := controller.productRepository.FetchByProduct(product)

			if err != nil {
				log.Printf(err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": product})
			c.Done()
			return
		}

		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)

		if err != nil {
			log.Printf("Invalid page number unsupported format %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid page number format! Page number should be an integer."})
			return
		}

		offers, pageCount, err := controller.offerRepository.FetchAll(page, OffersPerPage)

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
			"data":      offers,
			"pageCount": pageCount,
		})
		c.Done()
	}
}

func (controller *OfferController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := strconv.Atoi(id)

		if err != nil {
			log.Printf("Invalid offer ID %s", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request: OfferID should be an integer value"})
			return
		}

		offer, err := controller.offerRepository.FetchByID(id)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		if offer.ID == 0 {
			c.JSON(http.StatusOK, map[string]any{
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, map[string]any{
				"data": offer,
			})
		}
		c.Done()
	}
}

type Request struct {
	ProductSKU string       `json:"product_sku" binding:"required"`
	Offer      models.Offer `json:"offer" binding:"required"`
}

func (controller *OfferController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request = Request{}
		err := c.BindJSON(&request)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var product = models.Product{SKU: request.ProductSKU}
		product, err = controller.productRepository.FetchByProduct(product)

		var offer = models.Offer{}
		if product.ID == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "Fail Not implemented!"})
			return
			// TODO - Request product to ProductManagement
		} else {
			offer = request.Offer
		}
		offer.ProductID = product.ID
		offer, err = controller.offerRepository.Create(offer)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Success",
			"data":    offer,
		})
		c.Done()
	}
}

func (controller *OfferController) Update() gin.HandlerFunc { // todo - investigate changes on the variant when changing the offer.
	return func(c *gin.Context) {
		var offer = models.Offer{}
		err := c.BindJSON(&offer)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		offer, err = controller.offerRepository.Update(offer)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		time.Sleep(100)

		c.JSON(http.StatusOK, map[string]any{
			"message": "Offer updated successfully",
			"data":    offer,
		})
		c.Done()
	}
}

func (controller *OfferController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		offer, err := controller.offerRepository.FetchByID(id)
		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		offer.Deleted = true

		time.Sleep(100)
		offer, err = controller.offerRepository.Update(offer)
		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "Offer with id " + id + " Has been deleted successfully",
			"data":    offer,
		})
		c.Done()
	}
}
