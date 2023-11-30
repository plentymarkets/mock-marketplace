package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/request_binders/query_binders"
	"order-microservice/pkg/utils/logger"
)

func GetOrderById(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		parameters := query_binders.NewGetOrderByIdRequest()
		httpError := parameters.Bind(context)

		if httpError != nil {
			logger.Log("could not bind request", nil)
			context.AbortWithStatusJSON(httpError.Status, httpError.Message)
			return
		}

		orderRepository := repositories.NewOrderRepository(database)

		fields := map[string]string{
			"seller_id": parameters.SellerId,
			"id":        parameters.OrderId,
		}

		order, err := orderRepository.FindOneByFields(fields)

		if err != nil {
			logger.Log("Database transaction was not successful", err)
			context.AbortWithStatusJSON(http.StatusInternalServerError, "Database transaction was not successful")
			return
		}

		if order == nil {
			logger.Log("order not found", nil)
			context.AbortWithStatusJSON(http.StatusNotFound, "order not found")
			return
		}

		context.JSON(http.StatusOK, order)
		context.Done()
	}
}
