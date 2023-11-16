package request_binder

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
)

type CreateOrderRequest struct {
	OfferIds []int `json:"offerIds"`
}

func NewCreateOrderRequest() *CreateOrderRequest {
	return &CreateOrderRequest{}
}

func (request CreateOrderRequest) Bind(context *gin.Context) (*BindableRequest, *http_error.HttpError) {
	boundRequest, err := BindRequest(context, request)

	return boundRequest, err
}

func (request CreateOrderRequest) ValidateRequest() *http_error.HttpError {
	if len(request.OfferIds) == 0 {
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "no offer ids provided"}}
	}

	for _, id := range request.OfferIds {
		if id == 0 {
			return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "one or more provided offer ids is invalid"}}
		}
	}

	return nil
}
