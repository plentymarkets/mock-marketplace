package create_order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
)

type Request struct {
	OfferIds []int `json:"offerIds"`
}

func BindRequest(c *gin.Context) (*Request, *http_error.HttpError) {
	var request = &Request{}
	err := c.BindJSON(request)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid request body: %s", err.Error())}}
	}

	if len(request.OfferIds) == 0 {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "no offer ids provided"}}
	}

	for _, id := range request.OfferIds {
		if id == 0 {
			return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": "one or more provided offer ids is invalid"}}
		}
	}

	return request, nil
}
