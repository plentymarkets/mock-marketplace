package providers

import (
	"fmt"
	"order-microservice/pkg/routes/external-router"
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

func NewProductProvider() *Product {
	return &Product{}
}

func (product Product) Provide(route external_router.ExternalRoute, token string) (*Providable, error) {
	fetchRequest, err := FetchRequest(route, token)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve product: %s", err.Error())
	}

	return fetchRequest, err
}

func (product Product) Validate() error {
	return nil
}
