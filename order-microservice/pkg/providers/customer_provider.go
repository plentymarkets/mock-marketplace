package providers

import (
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http_error"
)

type Customer struct {
	ID              uint
	ShippingAddress string
	BillingAddress  string
}

func NewCustomerProvider() *Customer {
	return &Customer{}
}

func (customer *Customer) Provide(route *external_router.ExternalRoute, token *string) *http_error.HttpError {
	customer.ID = 1
	customer.ShippingAddress = "123 Main St, New York, NY 10001"
	customer.BillingAddress = "123 Main St, New York, NY 10001"

	if httpError := customer.ValidateProvided(); httpError != nil {
		return httpError
	}

	return nil
}

func (customer *Customer) ValidateProvided() *http_error.HttpError {
	return nil
}
