package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"offer-microservice/pkg/repositories"
	"strconv"
)

type offerController struct{}

func (controller *offerController) Createoffer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create offer
	}
}

func (controller *offerController) UpdateofferStatus(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.GetHeader("sellerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid seller id",
			})
			c.Abort()
			return
		}

		offerId, err := strconv.Atoi(c.GetHeader("offerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid offer id",
			})
			c.Abort()
			return
		}

		status := c.GetHeader("status")

		if status == "" {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid new status",
			})
			c.Abort()
			return
		}

		offerRepository := repositories.NewOfferRepository(databaseConnection)

		fields := map[string]string{
			"seller_id": strconv.Itoa(sellerId),
			"id":        strconv.Itoa(offerId),
		}

		offer, err := offerRepository.FindOneByFields(fields)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not retrieve offer",
			})
			c.Abort()
			return
		}

		if offer == nil {
			c.JSON(http.StatusNotFound, map[string]string{
				"error": "offer not found",
			})
			c.Abort()
			return
		}

		offer.Status = status
		transaction := offerRepository.Database.Save(&offer)

		if transaction.Error != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not update offer",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": "offer updated successfully",
		})

		c.Done()
	}
}

func (controller *offerController) Getoffers(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.GetHeader("sellerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid seller id",
			})
			c.Abort()
			return
		}

		offerRepository := repositories.NewOfferRepository(databaseConnection)
		offers, err := offerRepository.FindByField("seller_id", strconv.Itoa(sellerId), nil, nil)

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

func (controller *offerController) GetofferById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get offer by id
	}
}
