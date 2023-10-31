package providers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Authentication(url string, token string) (bool, error) {
	client := &http.Client{}

	data := map[string]string{"token": token}
	jsonData, err := json.Marshal(data)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}

	request.Header.Add("content-type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}
