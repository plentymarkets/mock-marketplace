package query_binders

import (
	"github.com/gin-gonic/gin"
	"order-microservice/pkg/utils/http_error"
)

type BindableQueryRequest interface {
	Bind(context *gin.Context) *http_error.HttpError
	ValidateQueryRequest() *http_error.HttpError
}
