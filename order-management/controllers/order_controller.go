package controllers

import (
	"github.com/gin-gonic/gin"
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

func (controller *OrderController) GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get orders
	}
}

func (controller *OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get order by id
	}
}
