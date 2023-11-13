package update_status

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
)

type Request struct {
	SellerId int    `json:"sellerId"`
	OrderId  int    `json:"orderId"`
	Status   string `json:"status"`
}

func BindRequest(c *gin.Context) (*Request, *http_error.HttpError) {
	var request = &Request{}
	err := c.BindJSON(request)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid request body: %s", err.Error())}}
	}

	if request.SellerId == 0 {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if request.OrderId == 0 {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid order id"}}
	}

	if request.Status == "" {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid status"}}
	}

	return request, nil
}
