package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/request_binders/query_binders"
	"order-microservice/pkg/utils/logger"
	"order-microservice/pkg/utils/string_conversion"
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

		sellerId, err := string_conversion.StringToUint(parameters.SellerId)

		if err != nil {
			logger.Log("could not convert seller id to integer", err)
			context.AbortWithStatusJSON(http.StatusBadRequest, "Seller id must be convertable to integer")
			return
		}

		OrderId, err := string_conversion.StringToUint(parameters.OrderId)

		if err != nil {
			logger.Log("could not convert order id to integer", err)
			context.AbortWithStatusJSON(http.StatusBadRequest, "Order id must be convertable to integer")
			return
		}

		order, err := orderRepository.FindOneBySellerAndByOrderId(sellerId, OrderId)

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
