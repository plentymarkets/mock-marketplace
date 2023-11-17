package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/services/order-service/update-status"
)

func UpdateOrderStatus(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		request, err := update_status.BindRequest(context)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		orderRepository := repositories.NewOrderRepository(database)
		result, err := update_status.FetchOrder(orderRepository, request)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		order, err := update_status.UpdateStatus(orderRepository, result.Order, request)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		context.JSON(http.StatusOK, order)
		context.Done()
	}
}
