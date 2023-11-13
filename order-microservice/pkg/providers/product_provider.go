package providers

import (
	"encoding/json"
	"io"
	"net/http"
	"order-microservice/pkg/routes/external_router"
)

type Product struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"userId"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	GTIN          string    `json:"gtin"`
	Categories    int       `json:"categories"`
	Manufacturers int       `json:"manufacturers"`
	Attributes    string    `json:"attributes"`
	Deleted       bool      `json:"deleted"`
	Variants      []Variant `json:"variants"`
}

type Variant struct {
	ID         uint   `json:"id"`
	ProductID  uint   `json:"productId"`
	UserID     uint   `json:"userId"`
	Name       string `json:"name"`
	GTIN       string `json:"gtin"`
	Attributes string `json:"attributes"`
	Deleted    bool   `json:"deleted"`
}

func FetchProduct(route external_router.ExternalRoute, token string) (*Product, error) {
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

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
