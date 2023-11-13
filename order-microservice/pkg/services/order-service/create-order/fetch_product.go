package create_order

import (
	"fmt"
	"net/http"
	"order-microservice/pkg/providers"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http-error"
)

type ProductResult struct {
	Product providers.Product
}

func FetchProduct(externalRouter external_router.ExternalRouter, gtin int, token string) (*ProductResult, *http_error.HttpError) {
	url := externalRouter.GetRoute("get-product-by-gtin", map[string]string{"gtin": fmt.Sprintf("%d", gtin)})
	product, err := providers.FetchProduct(url, token)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not retrieve product: %s", err.Error())}}
	}

	//TODO: Error handling like this currently not possible as there is no ID provided by the product service
	//if product.ID == 0 {
	//	return nil, &http_error.HttpError{Status: http.StatusNotFound, Message: map[string]string{"error": fmt.Sprintf("product not found: %s", err.Error())}}
	//}

	return &ProductResult{
		Product: *product,
	}, nil
}
