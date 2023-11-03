package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"offer-management/pkg/controllers"
	"offer-management/pkg/middlewares"
	"offer-management/pkg/repositories"
)

func Order(mariadb *gorm.DB, engine *gin.Engine) {

	orderRepository, _ := repositories.NewOrderRepository(mariadb)
	orderController := controllers.NewOrderController(orderRepository)

	order := engine.Group("/api/order").Use(middlewares.Authenticate())

	order.GET("/", orderController.GetAll())
	order.POST("/", orderController.Create())

}
