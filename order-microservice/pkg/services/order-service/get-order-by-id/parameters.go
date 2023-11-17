package get_order_by_id

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
	"strconv"
)

type Parameters struct {
	SellerId int
	OrderId  int
}

func InputParameters(c *gin.Context) (*Parameters, *http_error.HttpError) {
	sellerId, err := strconv.Atoi(c.MustGet("sellerId").(string))

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if sellerId == 0 {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid seller id")}}
	}

	orderId, err := strconv.Atoi(c.Param("orderId"))

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "invalid seller id"}}
	}

	if orderId == 0 {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid order id")}}
	}

	return &Parameters{
		SellerId: sellerId,
		OrderId:  orderId,
	}, nil
}
