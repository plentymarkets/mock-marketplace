package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-management/repositories"
	"strconv"
)

type OrderController struct{}

func (controller *OrderController) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create order
	}
}

func (controller *OrderController) UpdateOrderStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Update order status
	}
}

func (controller *OrderController) GetOrders(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.GetHeader("Authorization")
		sellerId, err := strconv.Atoi(c.Param("sellerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid seller id",
			})
		}
		//
		//if token == "" {
		//	c.JSON(http.StatusUnauthorized, map[string]string{
		//		"error": "Token is missing",
		//	})
		//}

		orderRepository := repositories.NewRepository(databaseConnection)
		orders := orderRepository.GetOrdersBySellerId(sellerId)

		c.JSON(http.StatusOK, orders)
	}
}

func (controller *OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get order by id
	}
}
