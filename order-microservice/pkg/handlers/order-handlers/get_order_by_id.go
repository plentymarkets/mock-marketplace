package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/services/order-service/get-order-by-id"
)

func GetOrderById(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		parameters, err := get_order_by_id.InputParameters(context)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		orderRepository := repositories.NewOrderRepository(database)
		orders, err := get_order_by_id.FetchOrderById(orderRepository, parameters)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		context.JSON(http.StatusOK, orders)
		context.Done()
	}
}
