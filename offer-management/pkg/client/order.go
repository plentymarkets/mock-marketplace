package client

import (
	"fmt"
	"net/http"
	"os"
)

type OrderClient struct {
	Endpoint string
	ApiToken string
}

func NewOrderClient(endpoint string, apiToken string) ProductContract {
	return OrderClient{
		Endpoint: endpoint,
		ApiToken: apiToken,
	}
}

func (client OrderClient) Get() (*http.Response, error) {
	request := Request{}
	request.URL = fmt.Sprintf("%s%s", os.Getenv("PRODUCTS_URL"), client.Endpoint)
	request.Token = client.ApiToken

	return GET(request)
}
