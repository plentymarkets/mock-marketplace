package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"offer-microservice/pkg/controllers"
	"offer-microservice/pkg/middlewares"
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
	router.RegisterofferRoutes()

	err := router.engine.Run(os.Getenv("GIN_PORT"))

	if err != nil {
		panic(err.Error())
	}
}

func (router Router) RegisterofferRoutes() {
	offerController := controllers.offerController{}

	routes := router.engine.Group("/offers").Use(middlewares.Authenticate(router.externalRouter.GetRoute("authenticationService")))
	routes.POST("/get", offerController.Getoffers(router.database))
	routes.POST("/update-status", offerController.UpdateofferStatus(router.database))
}
