package create_order

import (
	"fmt"
	"math/rand"
	"net/http"
	"order-microservice/pkg/models"
	"order-microservice/pkg/providers"
	"order-microservice/pkg/repositories"
	"order-microservice/pkg/routes/external_router"
	"order-microservice/pkg/utils/http-error"
	"time"
)

const (
	pending   = "pending"
	shipped   = "shipped"
	delivered = "delivered"
	cancelled = "cancelled"
)

func Create(repository repositories.OrderRepository, externalRouter external_router.ExternalRouter, request *Request, token string) *http_error.HttpError {
	customerResult, httpError := FetchCustomer()
	if httpError != nil {
		return httpError
	}

	orders := make(map[uint]*models.Order)

	for _, id := range request.OfferIds {
		offerResult, httpError := FetchOffer(externalRouter, id, token)
		if httpError != nil {
			return httpError
		}

		productResult, httpError := FetchProduct(externalRouter, offerResult.Offer.Gtin, token)
		if httpError != nil {
			return httpError
		}

		if orders[offerResult.Offer.SellerID] == nil {
			orders[offerResult.Offer.SellerID] = &models.Order{}
			addOrderData(orders[offerResult.Offer.SellerID], customerResult.Customer, offerResult.Offer)
		}

		addOrderItemData(orders[offerResult.Offer.SellerID], offerResult.Offer, productResult.Product)
		recalculateTotalSum(orders[offerResult.Offer.SellerID], offerResult.Offer)
	}

	for _, order := range orders {
		err := generateOrderNumber(order, repository)
		if err != nil {
			return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not generate order number: %s", err.Error())}}
		}

		err = repository.Create(order)
		if err != nil {
			return &http_error.HttpError{Status: http.StatusInternalServerError, Message: map[string]string{"error": fmt.Sprintf("could not create order: %s", err.Error())}}
		}
	}

	return nil
}

func recalculateTotalSum(order *models.Order, offer providers.Offer) {
	order.TotalSum += offer.Price
}

func addOrderItemData(order *models.Order, offer providers.Offer, product providers.Product) {
	var orderItem models.OrderItem

	orderItem.OrderID = order.ID
	orderItem.OfferID = offer.ID
	orderItem.GTIN = product.GTIN
	orderItem.Price = offer.Price
	orderItem.Quantity = 1      // placeholder as there is no real quantity
	orderItem.SKU = "123456789" // placeholder as there is no real SKU

	order.OrderItems = append(order.OrderItems, orderItem)
}

func addOrderData(order *models.Order, customer providers.Customer, offer providers.Offer) {
	order.CustomerID = customer.ID
	order.SellerID = offer.SellerID
	order.BillingAddress = customer.BillingAddress
	order.ShippingAddress = customer.ShippingAddress
	order.OrderDate = time.Now()
	order.Status = pending
	order.TotalSum = 0
}

func generateOrderNumber(order *models.Order, repository repositories.OrderRepository) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(10000) // generates a random integer up to 10000
	timestamp := time.Now().Unix()
	orderNumber := fmt.Sprintf("%d-%d", timestamp, randomNumber)

	result, err := repository.FindOneByField("order_number", orderNumber)

	if err != nil {
		return fmt.Errorf("could not generate order number: %s", err.Error())
	}

	if result != nil {
		err = generateOrderNumber(order, repository)

		if err != nil {
			return fmt.Errorf("could not generate order number: %s", err.Error())
		}
	}

	order.OrderNumber = orderNumber
	return nil
}
