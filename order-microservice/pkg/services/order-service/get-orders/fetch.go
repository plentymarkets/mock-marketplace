package get_orders

import (
	"fmt"
	"net/http"
	"order-microservice/pkg/models"
	"order-microservice/pkg/repositories"
	http_error "order-microservice/pkg/utils/http-error"
)

type Result struct {
	Orders []models.Order
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
