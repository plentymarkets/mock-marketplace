package request_binder

import (
	"github.com/gin-gonic/gin"
	"order-microservice/pkg/utils/http-error"
)

type BindableRequest interface {
	Bind(context *gin.Context) (*BindableRequest, *http_error.HttpError)
	ValidateRequest() *http_error.HttpError
}
