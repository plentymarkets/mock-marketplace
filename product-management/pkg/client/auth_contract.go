package client

import "net/http"

type AuthContract interface {
	Authenticate() (*http.Response, error)
	ValidateToken() (*http.Response, error)
}
