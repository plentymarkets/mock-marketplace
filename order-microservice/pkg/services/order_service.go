package services

import (
	"fmt"
	"net/http"
	"order-microservice/pkg/models"
	"order-microservice/pkg/providers"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
)

type OrderService struct {
	OrderRepository repositories.OrderRepository
	ExternalRouter  external_router.ExternalRouter
	apiKey          string
	token           string
}

func NewOrderService(orderRepository repositories.OrderRepository, externalRouter external_router.ExternalRouter, apiKey string, token string) *OrderService {
	return &OrderService{OrderRepository: orderRepository, ExternalRouter: externalRouter, apiKey: apiKey, token: token}
}

func (service *OrderService) CreateOrder(offerIds []int) *http_error.HttpError {
	customer := providers.NewCustomerProvider()
	httpError := customer.Provide(nil, nil)

	if httpError != nil {
		logger.Log("could not retrieve customer", nil)
		return httpError
	}

	order := models.Order{}
	order.AddOrderData(*customer)
	err := order.GenerateOrderNumber(service.OrderRepository)

	if err != nil {
		logger.Log("could not generate order number", err)
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": "could not generate order number"}}
	}

	for _, id := range offerIds {
		offer := providers.NewOfferProvider()
		route := service.ExternalRouter.GetRoute("get-offer-by-id", map[string]string{"id": fmt.Sprintf("%d", id)})
		httpError := offer.Provide(&route, &service.apiKey)

		if httpError != nil {
			logger.Log("could not retrieve offer", nil)
			return httpError
		}

		order.AddOrderItemData(*offer)
		order.RecalculateTotalSum(*offer)
	}

	err = service.OrderRepository.Create(&order)
	if err != nil {
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not create order: %s", err.Error())}}
	}

	return nil
}
