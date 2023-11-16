package request_binder

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
)

type UpdateOrderRequest struct {
	SellerId int    `json:"sellerId"`
	OrderId  int    `json:"orderId"`
	Status   string `json:"status"`
}

func NewUpdateOrderRequest() *UpdateOrderRequest {
	return &UpdateOrderRequest{}
}

func (request UpdateOrderRequest) Bind(context *gin.Context) (*BindableRequest, *http_error.HttpError) {
	boundRequest, err := BindRequest(context, request)

	return boundRequest, err
}

func (request UpdateOrderRequest) ValidateRequest() *http_error.HttpError {
	if request.SellerId == 0 {
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if request.OrderId == 0 {
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid order id"}}
	}

	if request.Status == "" {
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid status"}}
	}

	return nil
}
