package client

import (
	"fmt"
	"net/http"
	"os"
)

const ProductEndpoint = "product"

type ProductClient struct {
	Endpoint string
	ApiToken string
}

func NewProductClient(endpoint string, apiToken string) ProductClient {
	return ProductClient{
		Endpoint: endpoint,
		ApiToken: apiToken,
	}
}

func (client ProductClient) Get() (*http.Response, error) {

	request := Request{}
	request.URL = fmt.Sprintf("%s%s", os.Getenv("PRODUCTS_URL"), client.Endpoint)
	request.Token = client.ApiToken

	return GET(request)
}

func (client ProductClient) GetProduct(productSKU string) (*http.Response, error) {

	request := Request{}
	request.URL = fmt.Sprintf("%s/%s/%s", os.Getenv("PRODUCTS_URL"), ProductEndpoint, productSKU)
	request.Token = client.ApiToken

	return GET(request)
}
