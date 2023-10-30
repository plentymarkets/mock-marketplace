package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-management/controllers"
	"order-management/middlewares"
	"os"
)

func RegisterRoutes(databaseConnection *gorm.DB, externalRouter ExternalRouter) {
	orderController := controllers.OrderController{}

	router := gin.Default()

	routes := router.Group("/orders").Use(middlewares.Authenticate(externalRouter.GetRoute("authenticationService")))
	routes.POST("/get", orderController.GetOrders(databaseConnection))
	routes.POST("/update-status", orderController.UpdateOrderStatus(databaseConnection))

	err := router.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		panic(err.Error())
	}
}
