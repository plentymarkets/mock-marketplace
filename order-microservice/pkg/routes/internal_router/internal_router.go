package internal_router

import (
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

	err := internalRouter.Engine.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		panic(err.Error())
	}
}

func (internalRouter InternalRouter) RegisterOrderRoutes(database *gorm.DB) {
	routes := internalRouter.Engine.Group("/orders").Use(middlewares.Authenticate())
	routes.GET("/:sellerId", order_handlers.GetOrders(database))
	routes.PATCH("/status", order_handlers.UpdateOrderStatus(database))
	routes.POST("/", order_handlers.CreateOrder(database))
}
