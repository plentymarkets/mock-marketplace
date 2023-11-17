package internal_router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-microservice/pkg/handlers/order-handlers"
	"order-microservice/pkg/middlewares"
	"os"
)

type InternalRouter struct {
	Engine *gin.Engine
}

func NewInternalRouter() InternalRouter {
	var internalRouter InternalRouter
	internalRouter.Engine = gin.Default()

	return internalRouter
}

func (internalRouter InternalRouter) RegisterRoutes(database *gorm.DB) {
	internalRouter.RegisterOrderRoutes(database)

	address := fmt.Sprintf("%s:%s",
		os.Getenv("GIN_HOST"),
		os.Getenv("GIN_PORT"),
	)

	err := internalRouter.Engine.Run(address)

	if err != nil {
		panic(err.Error())
	}
}

func (internalRouter InternalRouter) RegisterOrderRoutes(database *gorm.DB) {
	routes := internalRouter.Engine.Group("/orders").Use(middlewares.Authenticate())
	routes.GET("/", order_handlers.GetOrders(database))
	routes.GET("/:orderId", order_handlers.GetOrderById(database))
	routes.PATCH("/status", order_handlers.UpdateOrderStatus(database))

	internalRoutes := internalRouter.Engine.Group("/internal/orders").Use(middlewares.AuthenticateApiKey())
	internalRoutes.POST("/", order_handlers.CreateOrder(database))
}
