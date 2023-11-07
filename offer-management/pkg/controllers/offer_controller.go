package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"offer-management/pkg/client"
	"offer-management/pkg/models"
	"offer-management/pkg/repositories"
	"strconv"
	"time"
)

const OffersPerPage = 10

type OfferController struct { // TODO - Move this after testing!
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

		gtin := c.DefaultQuery("gtin", "")
		if gtin != "" {

			var product = models.Product{GTIN: gtin}
			product, err := controller.productRepository.FetchByProduct(product)

			if err != nil {
				log.Printf(err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
				return
			}

			if product.ID == 0 {
				c.JSON(http.StatusOK, gin.H{"data": nil})
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
			c.JSON(http.StatusNotFound, gin.H{
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": offer,
			})
		}
		c.Done()
	}
}

type Request struct { // TODO - Mode to Models
	ProductGTIN string       `json:"product_gtin" binding:"required"`
	ProductSKU  string       `json:"product_sku"`
	Offer       models.Offer `json:"offer" binding:"required"`
}

func (controller *OfferController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var request = Request{}
		err := c.BindJSON(&request)

		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		var product = models.Product{GTIN: request.ProductGTIN}
		product, err = controller.productRepository.FetchByProduct(product)

		var offer = request.Offer
		offer.ID = 0

		if product.ID == 0 {

			apiToken := "ljHjuKSHDhduwhkHWUDHds8sd"

			productClient := client.NewProductClient("", apiToken)
			response, err := productClient.GetProduct("01234567890123")

			if err != nil {
				log.Print(err.Error())
				c.JSON(http.StatusNotFound, gin.H{"message": "Internal server Error"})
				return
			}

			if response.StatusCode != http.StatusOK {
				c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
				return
			}

			body, err := io.ReadAll(response.Body)
			err = json.Unmarshal(body, &product)

			if err != nil || product.GTIN == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error!"})
				return
			}

			product.Offers = append(product.Offers, offer)
			product, err = controller.productRepository.Create(product)

			c.JSON(http.StatusOK, product)
			return
		}

		offer.ProductID = product.ID
		offer, err = controller.offerRepository.Create(offer)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"data":    offer,
		})
		c.Done()
	}
}

func (controller *OfferController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var offer = models.Offer{}
		err := c.BindJSON(&offer)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "kldjshflkjashdfkl;ashjdfh"})
			return
		}

		offer, err = controller.offerRepository.Update(offer)

		if err != nil {
			log.Printf(err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
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

		c.JSON(http.StatusOK, gin.H{
			"message": "Offer with id " + id + " Has been deleted successfully",
			"data":    offer,
		})
		c.Done()
	}
}
