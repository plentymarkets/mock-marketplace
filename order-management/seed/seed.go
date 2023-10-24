package seed

import (
	"gorm.io/gorm"
	"math/rand"
	"order-management/models"
	"order-management/repositories"
	"time"
)

func Seed(databaseConnection *gorm.DB) {
	orderRepository := repositories.NewRepository(databaseConnection)
	order := generateOrder()
	orderRepository.DatabaseConnection.Create(&order)
}

func generateOrder() models.Order {
	order := models.Order{
		CustomerID:      rand.Int(),
		SellerID:        rand.Int(),
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
