package seed

import (
	"gorm.io/gorm"
	"math/rand"
	"order-microservice/pkg/models"
	"order-microservice/pkg/repositories"
	"time"
)

func Seed(databaseConnection *gorm.DB) error {
	orderRepository := repositories.NewOrderRepository(databaseConnection)
	isEmpty := checkIfTableIsEmpty(orderRepository)

	if isEmpty {
		order := generateOrder()
		transaction := orderRepository.Database.Create(&order)
		return transaction.Error
	}

	return nil
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
		orderItem := generateOrderItem(order)
		order.OrderItems = append(order.OrderItems, orderItem)
	}

	return order
}

func generateOrderItem(order models.Order) models.OrderItem {
	orderItem := models.OrderItem{
		OrderID:  order.ID,
		OfferID:  123,
		GTIN:     "1234567890123",
		Price:    rand.Float64(),
		Quantity: 1,
		SKU:      "1234567890123",
	}

	return orderItem
}

func checkIfTableIsEmpty(orderRepository repositories.OrderRepository) bool {
	var orders []models.Order
	orderRepository.Database.Find(&orders)
	return len(orders) == 0
}
