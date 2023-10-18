package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-management/controllers"
)

func RegisterRoutes(databaseConnection *gorm.DB) {
	orderController := controllers.OrderController{}

	router := gin.Default()

	routes := router.Group("/api/orders")
	// All these routes should be protected by authentication via .use(middlewares.Authenticate())
	routes.GET("/create", orderController.CreateOrder())
	routes.GET("/update/status/:id", orderController.UpdateOrderStatus())
	routes.GET("/get/:sellerId", orderController.GetOrders(databaseConnection))
	routes.GET("/get/:sellerId/:id", orderController.GetOrderById())

	router.Run()
}
