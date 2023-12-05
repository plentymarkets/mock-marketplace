package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/request_binders/body_binders"
	"order-microservice/pkg/utils/logger"
)

func UpdateOrderStatus(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		request := body_binders.NewUpdateOrderRequest()
		httpError := request.Bind(context)

		if httpError != nil {
			logger.Log("could not bind request", nil)
			context.AbortWithStatusJSON(httpError.Status, httpError.Message)
			return
		}

		orderRepository := repositories.NewOrderRepository(database)

		order, err := orderRepository.FindOneBySellerAndByOrderId(uint(request.SellerId), uint(request.OrderId))

		if err != nil {
			logger.Log("could not retrieve order", err)
			context.AbortWithStatusJSON(http.StatusInternalServerError, "Could not retrieve order")
			return
		}

		if order == nil {
			logger.Log("order not found", nil)
			context.AbortWithStatusJSON(http.StatusNotFound, "Order not found")
			return
		}

		order.Status = request.Status

		err = orderRepository.Update(order)

		if err != nil {
			logger.Log("could not update order", err)
			context.AbortWithStatusJSON(http.StatusInternalServerError, "Could not update order")
			return
		}

		context.JSON(http.StatusOK, order)
		context.Done()
	}
}
