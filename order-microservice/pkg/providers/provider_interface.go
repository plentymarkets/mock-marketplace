package providers

import "net/http"

type Providable interface {
	Provide(url string, token string) (*Providable, error)
	ValidateProvided(response *http.Response) error
}
