package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type AuthUserRequestBody struct {
	Username string `json:"email"`
	Password string `json:"password"`
	ApiToken string `json:"authenticationApiKey"`
}

func NewAuthUserRequest(username string, password string, apiToken string) ([]byte, error) {
	request := AuthUserRequestBody{
		Username: username,
		Password: password,
		ApiToken: apiToken,
	}

	return json.Marshal(request)
}

type AuthTokenRequestBody struct {
	ApiToken string `json:"token"`
}

func NewAuthTokenRequest(apiToken string) ([]byte, error) {
	request := AuthTokenRequestBody{
		ApiToken: apiToken,
	}

	return json.Marshal(request)
}

func Authenticate(username string, password string) (*http.Response, error) {

	body, err := NewAuthUserRequest(username, password, "test")
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	httpClient := &http.Client{}
	return httpClient.Post(
		"http://localhost:3001/user/token", // TODO - Remove hardcoding
		"application/json",
		bytes.NewBuffer(body),
	)
}

func ValidateToken(token string) (*http.Response, error) {

	body, err := NewAuthTokenRequest(token)

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	httpClient := &http.Client{} // TODO - Remove hardcoding
	return httpClient.Post(
		"http://localhost:3001/user/validation",
		"application/json",
		bytes.NewBuffer(body),
	)
}
