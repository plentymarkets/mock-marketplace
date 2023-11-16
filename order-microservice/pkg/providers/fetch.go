package providers

import (
	"io"
	"net/http"
	"order-microservice/pkg/routes/external-router"
)

func FetchRequest(route external_router.ExternalRoute, token string) (*Providable, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", route.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var providable Providable
	err = json.Unmarshal(body, &providable)
	if err != nil {
		return nil, err
	}

	return &providable, nil
}
