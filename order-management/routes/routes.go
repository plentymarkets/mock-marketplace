package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-management/controllers"
	"order-management/middlewares/authenticator"
	"os"
)

func RegisterRoutes(databaseConnection *gorm.DB, authenticator authenticator.Authenticator) {
	orderController := controllers.OrderController{}

	router := gin.Default()

	routes := router.Group("/orders").Use(authenticator.Authenticate())
	routes.POST("/get", orderController.GetOrders(databaseConnection))
	//routes.POST("/update-status", orderController.UpdateOrderStatus(databaseConnection))

	err := router.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		panic(err.Error())
	}
}
