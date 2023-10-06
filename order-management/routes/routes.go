package routes

import (
	"github.com/gin-gonic/gin"
	"order-management/controllers"
)

var orderController controllers.OrderController

func SetupRoutes() {
	router := gin.Default()

	routes := router.Group("/api/orders")
	// All these routes should be protected by authentication via .use(middlewares.Authenticate())
	routes.GET("/create", orderController.CreateOrder())
	routes.GET("/update/status/:id", orderController.UpdateOrderStatus())
	routes.GET("/get/:sellerId", orderController.GetOrders())
	routes.GET("/get/:sellerId/:id", orderController.GetOrderById())

	router.Run()
}
