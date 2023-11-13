package create_order

import (
	"order-microservice/pkg/providers"
	"order-microservice/pkg/utils/http-error"
)

type CustomerResult struct {
	Customer providers.Customer
}

func FetchCustomer() (*CustomerResult, *http_error.HttpError) {
	customer := providers.FetchCustomer()

	return &CustomerResult{
		Customer: customer,
	}, nil
}
