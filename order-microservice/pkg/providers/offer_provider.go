package providers

import (
	"encoding/json"
	"io"
	"net/http"
	"order-microservice/pkg/routes/external_router"
)

type Offer struct {
	ID       uint    `json:"id"`
	SellerID uint    `json:"sellerId"`
	GTIN     int     `json:"gtin"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func FetchOffer(route external_router.ExternalRoute, apiKey string) (*Offer, error) {
	client := &http.Client{}

	req, err := http.NewRequest(route.Method, route.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("apiKey", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var offer Offer
	err = json.Unmarshal(body, &offer)
	if err != nil {
		return nil, err
	}

	return &offer, nil
}
