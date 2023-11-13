package update_status

import (
	"fmt"
	"net/http"
	"order-microservice/pkg/models"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/utils/http-error"
)

func UpdateStatus(repository repositories.OrderRepository, order models.Order, request *Request) *http_error.HttpError {
	order.Status = request.Status
	err := repository.Update(&order)

	if err != nil {
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not update order: %s", err.Error())}}
	}

	return nil
}
