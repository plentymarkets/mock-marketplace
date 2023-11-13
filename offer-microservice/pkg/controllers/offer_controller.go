package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"offer-microservice/pkg/models"
	"offer-microservice/pkg/repositories"
	"strconv"
)

type OfferController struct{}

func (controller *OfferController) CreateOffer(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var offer = models.Offer{}
		err := c.BindJSON(&offer)

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
			c.Abort()
			return
		}

		offerRepository := repositories.NewOfferRepository(databaseConnection)
		offer, err = offerRepository.Create(offer)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not create offer",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, offer)
		c.Done()
	}
}

func (controller *OfferController) GetOffers(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId := c.Param("sellerId")

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid page",
			})
			c.Abort()
			return
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid limit",
			})
			c.Abort()
			return
		}

		offset := (page - 1) * limit

		offerRepository := repositories.NewOfferRepository(databaseConnection)
		offers, err := offerRepository.FindByField("seller_id", sellerId, &offset, &limit)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not retrieve offers",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, offers)
		c.Done()
	}
}

func (controller *OfferController) GetOfferById(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		offerId, err := strconv.Atoi(c.Param("offerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid offer id",
			})
			c.Abort()
			return
		}

		offerRepository := repositories.NewOfferRepository(databaseConnection)
		offer, err := offerRepository.FindOneById(offerId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not retrieve offer",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, offer)
		c.Done()
	}
}
