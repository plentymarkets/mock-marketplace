package request_binder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-microservice/pkg/utils/http-error"
)

func BindRequest(context *gin.Context, requestData BindableRequest) (*BindableRequest, *http_error.HttpError) {
	err := context.BindJSON(&requestData)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusBadRequest, Message: map[string]string{"error": fmt.Sprintf("invalid requestData body: %s", err.Error())}}
	}

	httpError := requestData.ValidateRequest()

	if httpError != nil {
		return nil, httpError
	}

	return &requestData, nil
}
