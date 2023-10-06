package seed

import (
	"gorm.io/gorm"
	"math/rand"
	"order-management/helper"
	"order-management/models"
	"order-management/repositories"
	"time"
)

var databaseConnection *gorm.DB

func init() {
	helper.LoadEnvVariables()
	databaseConnection = helper.GetDatabaseConnection()
}

func Seed() {
	orderRepository := repositories.SetupRepository(databaseConnection)
	order := generateOrder()
	orderRepository.CreateOrder(order)
}

func generateOrder() models.Order {
	order := models.Order{
		CustomerID:      rand.Int(),
		OrderNumber:     "123456789",
		OrderDate:       time.Now(),
		Status:          "Created",
		TotalSum:        rand.Float64(),
		BillingAddress:  "123 Main Street, Apt 4B, Imaginary City, Utopia, 54321",
		ShippingAddress: "123 Main Street, Apt 4B, Imaginary City, Utopia, 54321",
	}

	for i := 0; i < 3; i++ {
		orderItem := models.OrderItem{
			OrderID:  order.ID,
			OfferID:  123,
			GTIN:     "1234567890123",
			Price:    rand.Float64(),
			Quantity: 1,
			SKU:      "1234567890123",
		}

		order.OrderItems = append(order.OrderItems, orderItem)
	}

	return order
}
