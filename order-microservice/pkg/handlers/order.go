package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/request-binder"
	"order-microservice/pkg/routes/external-router"
	"order-microservice/pkg/services/order-service/create-order"
	"order-microservice/pkg/services/order-service/get-orders"
	"order-microservice/pkg/services/order-service/update-status"
	http_error "order-microservice/pkg/utils/http-error"
)

func CreateOrder(databaseConnection *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		bindableRequest := request_binder.NewCreateOrderRequest()
		request, err := bindableRequest.Bind(context)

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

func GetOrders(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		parameters, err := get_orders.InputParameters(context)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		orderRepository := repositories.NewOrderRepository(database)
		offset := calculateOffset(parameters.Page, parameters.Limit)
		orders, err := orderRepository.FindByField("seller_id", parameters.SellerId, &offset, &parameters.Limit)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("could not retrieve orders: %s", err.Error())})
			return
		}

		context.JSON(http.StatusOK, orders)
		context.Done()
	}
}

func UpdateOrderStatus(database *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		bindableRequest := request_binder.NewUpdateOrderRequest()
		request, err := bindableRequest.Bind(context)
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

		err = update_status.UpdateStatus(orderRepository, result.Order, request)
		if err != nil {
			context.AbortWithStatusJSON(err.Status, err.Message)
			return
		}

		context.JSON(http.StatusOK, map[string]string{})
		context.Done()
	}
}

func FetchOrders(repository repositories.OrderRepository, parameters *Parameters) (*Result, *http_error.HttpError) {
	offset := calculateOffset(parameters.Page, parameters.Limit)

	orders, err := repository.FindByField("seller_id", parameters.SellerId, &offset, &parameters.Limit)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("could not retrieve orders: %s", err.Error())}}
	}

	return &Result{
		Orders: *orders,
	}, nil
}

func calculateOffset(page int, limit int) int {
	offset := (page - 1) * limit
	return offset
}
