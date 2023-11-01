package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-microservice/pkg/controllers"
	"order-microservice/pkg/middlewares"
	"os"
)

type Router struct {
	engine         *gin.Engine
	externalRouter *ExternalRouter
	database       *gorm.DB
}

func (router Router) NewRouter(databaseConnection *gorm.DB) Router {
	router.engine = gin.Default()

	var externalRouter ExternalRouter
	externalRouter = externalRouter.NewExternalRouter()
	router.externalRouter = &externalRouter

	router.database = databaseConnection

	return router
}

func (router Router) RegisterRoutes() {
	router.RegisterOrderRoutes()

	err := router.engine.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		panic(err.Error())
	}
}

func (router Router) RegisterOrderRoutes() {
	orderController := controllers.OrderController{}

	routes := router.engine.Group("/orders").Use(middlewares.Authenticate(router.externalRouter.GetRoute("authenticationService")))
	routes.POST("/get", orderController.GetOrders(router.database))
	routes.POST("/update-status", orderController.UpdateOrderStatus(router.database))
}
