package order_handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/request_binders/query_binders"
	"order-microservice/pkg/utils/logger"
	. "order-microservice/pkg/utils/string_conversion"
)

func GetOrders(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		parameters := query_binders.NewGetOrdersRequest()
		httpError := parameters.Bind(context)

		if httpError != nil {
			logger.Log("could not bind request", nil)
			context.AbortWithStatusJSON(httpError.Status, httpError.Message)
			return
		}

		page, err := StringToInt(parameters.Page)
		if err != nil {
			logger.Log("could not convert page to integer", err)
			context.AbortWithStatusJSON(http.StatusBadRequest, "Seller id must be an integer")
			return
		}

		limit, err := StringToInt(parameters.Limit)
		if err != nil {
			logger.Log("could not convert limit to integer", err)
			context.AbortWithStatusJSON(http.StatusBadRequest, "Page must be an integer")
			return
		}

		offset := calculateOffset(page, limit)

		orderRepository := repositories.NewOrderRepository(database)
		orders, err := orderRepository.FindByField("seller_id", parameters.SellerId, &offset, &limit)

		if err != nil {
			logger.Log("Database transaction was not successful", err)
			context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		if orders == nil {
			logger.Log("No orders found", nil)
			context.AbortWithStatusJSON(http.StatusNotFound, "No orders found")
			return
		}

		context.JSON(http.StatusOK, orders)
		context.Done()
	}
}

func calculateOffset(page int, limit int) int {
	offset := (page - 1) * limit
	return offset
}
