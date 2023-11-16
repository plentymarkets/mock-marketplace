package providers

import (
	"fmt"
	"order-microservice/pkg/routes/external-router"
)

type Offer struct {
	ID       uint    `json:"id"`
	SellerID uint    `json:"sellerId"`
	Gtin     int     `json:"gtin"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func NewOfferProvider() *Offer {
	return &Offer{}
}

func (offer Offer) Provide(route external_router.ExternalRoute, token string) (*Providable, error) {
	fetchRequest, err := FetchRequest(route, token)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve offer: %s", err.Error())
	}

	return fetchRequest, err
}

func (offer Offer) Validate() error {
	if offer.ID == 0 {
		return fmt.Errorf("offer not found")
	}

	return nil
}
