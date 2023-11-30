package body_binders

import (
	"github.com/gin-gonic/gin"
	"order-microservice/pkg/utils/http_error"
)

type BindableBodyRequest interface {
	Bind(context *gin.Context) *http_error.HttpError
	ValidateBodyRequest() *http_error.HttpError
}
