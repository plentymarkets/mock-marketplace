package body_binders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
	"order-microservice/pkg/utils/string_conversion"
)

type UpdateOrderRequest struct {
	SellerId int
	OrderId  int    `json:"orderId"`
	Status   string `json:"status"`
}

func NewUpdateOrderRequest() *UpdateOrderRequest {
	return &UpdateOrderRequest{}
}

func (request *UpdateOrderRequest) Bind(context *gin.Context) *http_error.HttpError {
	sellerId := context.GetString("sellerId")
	err := context.BindJSON(&request)

	request.SellerId, err = string_conversion.StringToInt(sellerId)

	if err != nil {
		logger.Log("invalid requestData body", err)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid requestData body: %s"}}
	}

	if httpError := request.ValidateBodyRequest(); httpError != nil {
		return httpError
	}

	return nil
}

func (request *UpdateOrderRequest) ValidateBodyRequest() *http_error.HttpError {
	if request.SellerId == 0 {
		logger.Log("invalid seller id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if request.OrderId == 0 {
		logger.Log("invalid order id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid order id"}}
	}

	if request.Status == "" {
		logger.Log("invalid status", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid status"}}
	}

	return nil
}
