package providers

import (
	"io"
	"net/http"
	"order-microservice/pkg/routes/external_router"
)

func FetchRequest(route external_router.ExternalRoute, header map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(route.Method, route.Url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
