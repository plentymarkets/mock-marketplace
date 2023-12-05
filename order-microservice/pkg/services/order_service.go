package services

import (
	"fmt"
	"math/rand"
	"net/http"
	"order-microservice/pkg/models"
	"order-microservice/pkg/providers"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http_error"
	"order-microservice/pkg/utils/logger"
	"time"
)

type OrderService struct {
	OrderRepository repositories.OrderRepository
	ExternalRouter  external_router.ExternalRouter
	Order           models.Order
	ApiKey          string
	Token           string
}

func NewOrderService(orderRepository repositories.OrderRepository, externalRouter external_router.ExternalRouter, apiKey string, token string) *OrderService {
	return &OrderService{OrderRepository: orderRepository, ExternalRouter: externalRouter, ApiKey: apiKey, Token: token}
}

func (Service *OrderService) CreateOrder(offerIds []int) *http_error.HttpError {
	customer := providers.NewCustomerProvider()
	httpError := customer.Provide(nil, nil)

	if httpError != nil {
		logger.Log("could not retrieve customer", nil)
		return httpError
	}

	Service.AddOrderData(*customer)
	err := Service.GenerateOrderNumber(Service.OrderRepository)

	if err != nil {
		logger.Log("could not generate order number", err)
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": "could not generate order number"}}
	}

	for _, id := range offerIds {
		offer := providers.NewOfferProvider()
		route := Service.ExternalRouter.GetRoute("get-offer-by-id", map[string]string{"offerId": fmt.Sprintf("%d", id)})
		httpError := offer.Provide(&route, &Service.ApiKey)

		if httpError != nil {
			logger.Log("could not retrieve offer", nil)
			return httpError
		}

		Service.AddOrderItemData(*offer)
		Service.RecalculateTotalSum(*offer)
	}

	err = Service.OrderRepository.Create(&Service.Order)
	if err != nil {
		logger.Log("could not create order", err)
		return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": "could not create order: %s"}}
	}

	return nil
}

func (Service *OrderService) RecalculateTotalSum(offer providers.Offer) {
	Service.Order.TotalSum += offer.Price
}

func (Service *OrderService) AddOrderItemData(offer providers.Offer) {
	var orderItem models.OrderItem

	orderItem.OrderID = Service.Order.ID
	orderItem.SellerID = offer.SellerID
	orderItem.OfferID = offer.ID
	orderItem.GTIN = offer.Gtin
	orderItem.Price = offer.Price
	orderItem.Quantity = 1      // placeholder as there is no real quantity
	orderItem.SKU = "123456789" // placeholder as there is no real SKU

	Service.Order.OrderItems = append(Service.Order.OrderItems, orderItem)
}

func (Service *OrderService) AddOrderData(customer providers.Customer) {
	Service.Order.CustomerID = customer.ID
	Service.Order.BillingAddress = customer.BillingAddress
	Service.Order.ShippingAddress = customer.ShippingAddress
	Service.Order.OrderDate = time.Now()
	Service.Order.Status = "pending"
	Service.Order.TotalSum = 0
}

func (Service *OrderService) GenerateOrderNumber(repository repositories.OrderRepository) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(10000)
	timestamp := time.Now().Unix()
	orderNumber := fmt.Sprintf("%d-%d", timestamp, randomNumber)

	result, err := repository.FindOneByField("order_number", orderNumber)

	if err != nil {
		return fmt.Errorf("could not generate order number: %s", err.Error())
	}

	if result != nil {
		err = Service.GenerateOrderNumber(repository)

		if err != nil {
			return fmt.Errorf("could not generate order number: %s", err.Error())
		}
	}

	Service.Order.OrderNumber = orderNumber
	return nil
}
