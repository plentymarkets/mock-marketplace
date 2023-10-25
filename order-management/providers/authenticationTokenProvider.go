package providers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

// FetchToken
// TODO: Should email and password be sent via Header?
func FetchToken(url string, email string, password string, authenticationApiKey string) (*Token, error) {
	client := &http.Client{}

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("authenticationApiKey", authenticationApiKey)
	request.Header.Add("email", email)
	request.Header.Add("password", password)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
