package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/request_binders/body_binders"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/services"
	"order-microservice/pkg/utils/logger"
	"os"
)

func CreateOrder(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		request := body_binders.NewCreateOrderRequest()
		httpError := request.Bind(context)

		if httpError != nil {
			logger.Log("could not bind request", nil)
			context.AbortWithStatusJSON(httpError.Status, httpError.Message)
			return
		}

		externalRouter := external_router.NewExternalRouter()

		orderRepository := repositories.NewOrderRepository(databaseConnection)

		orderService := services.NewOrderService(orderRepository, externalRouter, os.Getenv("API_KEY"), context.GetHeader("sellerId"))
		httpErr := orderService.CreateOrder(request.OfferIds)

		if httpErr != nil {
			logger.Log("could not create order", nil)
			context.AbortWithStatusJSON(httpErr.Status, httpErr.Message)
			return
		}

		context.JSON(http.StatusOK, map[string]string{"message": "Order created successfully"})
		context.Done()
	}
}
