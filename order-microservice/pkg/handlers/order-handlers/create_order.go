package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/services/order-service/create-order"
)

func CreateOrder(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		request, err := create_order.BindRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		externalRouter := external_router.NewExternalRouter()

		orderRepository := repositories.NewOrderRepository(databaseConnection)

		err = create_order.Create(orderRepository, externalRouter, request, context.GetHeader("token"))
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		context.JSON(http.StatusOK, map[string]string{"message": "Order created successfully"})
		context.Done()
	}
}
