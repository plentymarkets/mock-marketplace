package client

import (
	"fmt"
	"net/http"
	"os"
)

type ProductClient struct {
	Endpoint string
	ApiToken string
}

func NewProductClient(endpoint string, apiToken string) ProductContract {
	return ProductClient{
		Endpoint: endpoint,
		ApiToken: apiToken,
	}
}

func (client ProductClient) GetRequest() (*http.Response, error) {

	request := Request{}
	request.URL = fmt.Sprintf("%s%s", os.Getenv("PRODUCTS_URL"), client.Endpoint)
	request.Token = client.ApiToken

	return GET(request)
}
