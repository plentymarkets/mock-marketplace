package update_status

import (
	"fmt"
	"net/http"
	"order-microservice/pkg/models"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/utils/http-error"
	"strconv"
)

type Result struct {
	Order models.Order
}

func FetchOrder(repository repositories.OrderRepository, request *Request) (*Result, *http_error.HttpError) {
	fields := map[string]string{
		"seller_id": strconv.Itoa(request.SellerId),
		"id":        strconv.Itoa(request.OrderId),
	}

	order, err := repository.FindOneByFields(fields)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not retrieve order: %s", err.Error())}}
	}

	if order.ID == 0 {
		return nil, &http_error.HttpError{Status: http.StatusNotFound, Message: map[string]string{"error": fmt.Sprintf("Order not found: %s", err.Error())}}
	}

	return &Result{
		Order: *order,
	}, nil
}
