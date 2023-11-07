package client

import "net/http"

type ProductContract interface {
	Get() (*http.Response, error)
}
