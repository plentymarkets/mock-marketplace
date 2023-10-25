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
		sellerId, err := strconv.Atoi(c.GetHeader("sellerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid seller id",
			})
		}

		orderRepository := repositories.NewOrderRepository(databaseConnection)
		orders, err := orderRepository.FindByField("seller_id", strconv.Itoa(sellerId), nil, nil)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not retrieve orders",
			})
		}

		c.JSON(http.StatusOK, orders)
	}
}

func (controller *OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get order by id
	}
}
