package client

import (
	"bytes"
	"io"
	"net/http"
)

const ContentType = "application/json"

type Request struct {
	URL     string `json:"url,omitempty"`
	Token   string `json:"token,omitempty"`
	Headers string `json:"headers,omitempty"`
	Body    []byte `json:"body,omitempty"`
}

func GET(request Request) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, request.URL, bytes.NewBuffer(request.Body))
	if err != nil {
		// TODO - handle error
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", request.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// TODO - handle error
	}
	return resp, nil
}

func POST(request Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Post(request.URL, ContentType, bytes.NewBuffer(request.Body))
}

func PUT(request Request) {
	body := request.Body
	req, err := http.NewRequest(http.MethodPut, request.URL, bytes.NewBuffer(body))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}
