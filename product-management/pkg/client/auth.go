package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AuthenticationResponse struct {
	SellerId string `json:"sellerId"`
}

func AuthenticationRequest(token string) (*AuthenticationResponse, error) {
	client := &http.Client{}

	authURL := fmt.Sprintf("%s%s",
		os.Getenv("AUTHENTICATOR_MICROSERVICE_URL"),
		"/user/validation",
	)

	req, err := http.NewRequest("GET", authURL, nil)
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

	var authentication AuthenticationResponse
	err = json.Unmarshal(body, &authentication)
	if err != nil {
		return nil, err
	}

	return &authentication, nil
}
