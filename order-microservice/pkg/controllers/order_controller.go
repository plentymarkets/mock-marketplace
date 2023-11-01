package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"strconv"
)

type OrderController struct{}

func (controller *OrderController) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create order
	}
}

func (controller *OrderController) UpdateOrderStatus(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.GetHeader("sellerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid seller id",
			})
			c.Abort()
			return
		}

		orderId, err := strconv.Atoi(c.GetHeader("orderId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid order id",
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

		orderRepository := repositories.NewOrderRepository(databaseConnection)

		fields := map[string]string{
			"seller_id": strconv.Itoa(sellerId),
			"id":        strconv.Itoa(orderId),
		}

		order, err := orderRepository.FindOneByFields(fields)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not retrieve order",
			})
			c.Abort()
			return
		}

		if order == nil {
			c.JSON(http.StatusNotFound, map[string]string{
				"error": "Order not found",
			})
			c.Abort()
			return
		}

		order.Status = status
		transaction := orderRepository.Database.Save(&order)

		if transaction.Error != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not update order",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"message": "Order updated successfully",
		})

		c.Done()
	}
}

func (controller *OrderController) GetOrders(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sellerId, err := strconv.Atoi(c.GetHeader("sellerId"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid seller id",
			})
			c.Abort()
			return
		}

		orderRepository := repositories.NewOrderRepository(databaseConnection)
		orders, err := orderRepository.FindByField("seller_id", strconv.Itoa(sellerId), nil, nil)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not retrieve orders",
			})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, orders)
		c.Done()
	}
}

func (controller *OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get order by id
	}
}
