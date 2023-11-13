package get_orders

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
	"strconv"
)

type Parameters struct {
	SellerId string
	Page     int
	Limit    int
}

func InputParameters(c *gin.Context) (*Parameters, *http_error.HttpError) {
	sellerId := c.Param("sellerId")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid page: %s", err.Error())}}
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid limit: %s", err.Error())}}
	}

	return &Parameters{
		SellerId: sellerId,
		Page:     page,
		Limit:    limit,
	}, nil
}
