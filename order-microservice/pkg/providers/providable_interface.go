package providers

import (
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http_error"
)

type Providable interface {
	Provide(route *external_router.ExternalRoute, token *string) *http_error.HttpError
	ValidateProvided() *http_error.HttpError
}
