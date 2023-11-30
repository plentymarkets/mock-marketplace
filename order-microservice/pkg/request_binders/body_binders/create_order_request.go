package body_binders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
)

type CreateOrderRequest struct {
	OfferIds []int `json:"offerIds"`
}

func NewCreateOrderRequest() *CreateOrderRequest {
	return &CreateOrderRequest{}
}

func (request CreateOrderRequest) Bind(context *gin.Context) *http_error.HttpError {
	err := context.BindJSON(&request)

	if err != nil {
		logger.Log("invalid requestData body", err)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid requestData body"}}
	}

	if httpError := request.ValidateBodyRequest(); httpError != nil {
		return httpError
	}

	return nil
}

func (request CreateOrderRequest) ValidateBodyRequest() *http_error.HttpError {
	if len(request.OfferIds) == 0 {
		logger.Log("no offer ids provided", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "no offer ids provided"}}
	}

	for _, id := range request.OfferIds {
		if id == 0 {
			logger.Log("one or more provided offer ids is invalid", nil)
			return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "one or more provided offer ids is invalid"}}
		}
	}

	return nil
}
