package create_order

import (
	"fmt"
	"net/http"
	"order-microservice/pkg/providers"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http-error"
)

type OfferResult struct {
	Offer providers.Offer
}

func FetchOffer(externalRouter external_router.ExternalRouter, id int, token string) (*OfferResult, *http_error.HttpError) {
	url := externalRouter.GetRoute("get-offer-by-id", map[string]string{"offerId": fmt.Sprintf("%d", id)})
	offer, err := providers.FetchOffer(url, token)

	if err != nil {
		return nil, &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not retrieve offer: %s", err.Error())}}
	}

	if offer.ID == 0 {
		return nil, &http_error.HttpError{Status: http.StatusNotFound, Message: map[string]string{"error": "offer not found"}}
	}

	return &OfferResult{
		Offer: *offer,
	}, nil
}
