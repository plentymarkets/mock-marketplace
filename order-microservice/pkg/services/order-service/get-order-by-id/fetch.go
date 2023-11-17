package get_order_by_id

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

func FetchOrderById(repository repositories.OrderRepository, parameters *Parameters) (*Result, *http_error.HttpError) {
	fields := map[string]string{
		"seller_id": strconv.Itoa(parameters.SellerId),
		"id":        strconv.Itoa(parameters.OrderId),
	}

	order, err := repository.FindOneByFields(fields)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("could not retrieve orders: %s", err.Error())}}
	}

	return &Result{
		Order: *order,
	}, nil
}
