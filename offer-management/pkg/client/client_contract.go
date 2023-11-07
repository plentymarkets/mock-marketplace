package client

import "net/http"

type ProductContract interface {
	GetRequest() (*http.Response, error)
}
