package client

import "net/http"

func GET(url string) (*http.Response, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Token", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM5OTQ4NzB9.iuwberVnb1dcvnwCdqDPF1m25v4Dp-tgGpcW08zYjjc"`)

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return response, nil
}
