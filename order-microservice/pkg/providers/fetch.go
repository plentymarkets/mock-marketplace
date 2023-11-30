package providers

import (
	"encoding/json"
	"io"
	"net/http"
	"order-microservice/pkg/routes/external_router"
)

func FetchRequest(providable Providable, route external_router.ExternalRoute, token string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", route.Url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("token", token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &providable)
	if err != nil {
		return err
	}

	return nil
}
