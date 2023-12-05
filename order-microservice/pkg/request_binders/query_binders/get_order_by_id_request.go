package query_binders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
)

type CreateOrderByIdRequest struct {
	SellerId string
	OrderId  string
}

func NewGetOrderByIdRequest() *CreateOrderByIdRequest {
	return &CreateOrderByIdRequest{}
}

func (request *CreateOrderByIdRequest) Bind(c *gin.Context) *http_error.HttpError {
	sellerId := c.GetString("sellerId")
	orderId := c.Param("orderId")

	request.SellerId = sellerId
	request.OrderId = orderId

	if httpError := request.ValidateQueryRequest(); httpError != nil {
		return httpError
	}

	return nil
}

func (request *CreateOrderByIdRequest) ValidateQueryRequest() *http_error.HttpError {
	if request.SellerId == "" {
		logger.Log("invalid seller id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if request.OrderId == "" {
		logger.Log("invalid order id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid order id"}}
	}

	return nil
}
