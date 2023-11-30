package query_binders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
)

type GetOrdersRequest struct {
	SellerId string
	Page     string
	Limit    string
}

func NewGetOrdersRequest() *GetOrdersRequest {
	return &GetOrdersRequest{}
}

func (request *GetOrdersRequest) Bind(context *gin.Context) *http_error.HttpError {
	sellerId := context.GetString("sellerId")
	page := context.DefaultQuery("page", "1")
	limit := context.DefaultQuery("limit", "10")

	request.SellerId = sellerId
	request.Page = page
	request.Limit = limit

	if httpError := request.ValidateQueryRequest(); httpError != nil {
		return httpError
	}

	return nil
}

func (request *GetOrdersRequest) ValidateQueryRequest() *http_error.HttpError {
	if request.SellerId == "" {
		logger.Log("invalid seller id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if request.Page == "" {
		logger.Log("invalid order id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid order id"}}
	}

	if request.Limit == "" {
		logger.Log("invalid order id", nil)
		return &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid order id"}}
	}

	return nil
}
