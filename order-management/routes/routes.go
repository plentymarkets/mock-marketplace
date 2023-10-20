package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-management/controllers"
	"order-management/middlewares/authenticator"
)

func RegisterRoutes(databaseConnection *gorm.DB, authenticator authenticator.AuthenticatorInterface) {
	orderController := controllers.OrderController{}

	router := gin.Default()

	routes := router.Group("/api/orders").Use(authenticator.Authenticate())
	// All these routes should be protected by authentication via .use(middlewares.Authenticate())
	//routes.GET("/create", orderController.CreateOrder())
	//routes.GET("/update/status/:id", orderController.UpdateOrderStatus())
	routes.GET("/get/:sellerId", orderController.GetOrders(databaseConnection))
	//routes.GET("/get/:sellerId/:id", orderController.GetOrderById())

	err := router.Run()

	if err != nil {
		panic(err.Error())
	}
}
