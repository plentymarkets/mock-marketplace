package client

import "net/http"

func GET(url string, token string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Token", token)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return response, nil
}
