package routes

import (
	"fmt"
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
	router.RegisterOfferRoutes()

	address := fmt.Sprintf("%s:%s",
		os.Getenv("GIN_HOST"),
		os.Getenv("GIN_PORT"),
	)

	err := router.engine.Run(address)

	if err != nil {
		panic(err.Error())
	}
}

func (router Router) RegisterOfferRoutes() {
	offerController := controllers.OfferController{}

	routes := router.engine.Group("/offers").Use(middlewares.AuthenticateToken(router.externalRouter.GetRoute("authenticationService")))
	routes.POST("/", offerController.CreateOffer(router.database))
	routes.GET("/", offerController.GetSellersOffers(router.database))
	routes.GET("/:offerId", offerController.GetSellersOfferById(router.database))

	internalRoutes := router.engine.Group("/internal/offers").Use(middlewares.AuthenticateSecret())
	internalRoutes.GET("/:offerId", offerController.GetOfferById(router.database))
}
