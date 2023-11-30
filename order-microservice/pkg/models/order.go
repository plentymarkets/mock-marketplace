package models

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"order-microservice/pkg/providers"
	"order-microservice/pkg/repositories"
	"time"
)

const (
	pending   = "pending"
	shipped   = "shipped"
	delivered = "delivered"
	cancelled = "cancelled"
)

type Order struct {
	ID              uint      `gorm:"primaryKey"`
	CustomerID      uint      `gorm:"type:uint"`
	OrderNumber     string    `gorm:"type:varchar(255)"`
	OrderDate       time.Time `gorm:"type:datetime"`
	Status          string    `gorm:"type:varchar(50)"`
	TotalSum        float64   `gorm:"type:decimal(10,2)"`
	BillingAddress  string    `gorm:"type:varchar(255)"`
	ShippingAddress string    `gorm:"type:varchar(255)"`
	OrderItems      []OrderItem
}

func (order Order) Migrate(databaseConnection *gorm.DB) error {
	return databaseConnection.AutoMigrate(order)
}

func (order Order) RecalculateTotalSum(offer providers.Offer) {
	order.TotalSum += offer.Price
}

func (order Order) AddOrderItemData(offer providers.Offer) {
	var orderItem OrderItem

	orderItem.OrderID = order.ID
	orderItem.SellerID = offer.SellerID
	orderItem.OfferID = offer.ID
	orderItem.GTIN = offer.Gtin
	orderItem.Price = offer.Price
	orderItem.Quantity = 1      // placeholder as there is no real quantity
	orderItem.SKU = "123456789" // placeholder as there is no real SKU

	order.OrderItems = append(order.OrderItems, orderItem)
}

func (order Order) AddOrderData(customer providers.Customer) {
	order.CustomerID = customer.ID
	order.BillingAddress = customer.BillingAddress
	order.ShippingAddress = customer.ShippingAddress
	order.OrderDate = time.Now()
	order.Status = pending
	order.TotalSum = 0
}

func (order Order) GenerateOrderNumber(repository repositories.OrderRepository) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(10000)
	timestamp := time.Now().Unix()
	orderNumber := fmt.Sprintf("%d-%d", timestamp, randomNumber)

	result, err := repository.FindOneByField("order_number", orderNumber)

	if err != nil {
		return fmt.Errorf("could not generate order number: %s", err.Error())
	}

	if result != nil {
		err = order.GenerateOrderNumber(repository)

		if err != nil {
			return fmt.Errorf("could not generate order number: %s", err.Error())
		}
	}

	order.OrderNumber = orderNumber
	return nil
}
