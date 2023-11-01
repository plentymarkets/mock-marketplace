package client

import (
	"bytes"
	"net/http"
)

const ContentType = "application/json"

type Request struct {
	URL      string `json:"url,omitempty"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	JsonBody string `json:"body,omitempty"`
}

func GET(request Request) (*http.Response, error) {
	body := []byte(request.JsonBody)
	client := &http.Client{}
	return client.Post(request.URL, ContentType, bytes.NewBuffer(body))
}

func POST(request Request) (*http.Response, error) {
	body := []byte(request.JsonBody)
	client := &http.Client{}
	return client.Post(request.URL, ContentType, bytes.NewBuffer(body))
}

func PUT(request Request) {
	body := []byte(request.JsonBody)
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
	defer resp.Body.Close()
}

func DEL() {

}
