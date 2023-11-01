package client

import (
	"fmt"
	"net/http"
	"os"
)

type AuthClient struct {
	Endpoint string
	ApiToken string
	Username string
	Password string
}

func NewAuthClient(endpoint string, apiToken string, username string, password string) AuthClient {
	return AuthClient{
		Endpoint: endpoint,
		ApiToken: apiToken,
		Username: username,
		Password: password,
	}
}

func (client AuthClient) Authenticate() (*http.Response, error) {

	request := Request{}
	request.URL = fmt.Sprintf("%s%s", os.Getenv("AUTH_URL"), "/")
	request.Username = client.Username
	request.Password = client.Password

	return GET(request)
}

func (client AuthClient) ValidateToken() (*http.Response, error) {

	request := Request{}
	request.URL = fmt.Sprintf("%s%s", os.Getenv("AUTH_URL"), client.Endpoint)
	request.Token = client.ApiToken

	return GET(request)
}

func (client AuthClient) GetRequest() (*http.Response, error) {

	request := Request{}
	request.URL = fmt.Sprintf("%s%s", os.Getenv("PRODUCTS_URL"), client.Endpoint)
	request.Token = client.ApiToken

	return GET(request)
}
