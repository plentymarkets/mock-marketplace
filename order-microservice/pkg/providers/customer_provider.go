package providers

type Customer struct {
	ID              uint
	ShippingAddress string
	BillingAddress  string
}

func FetchCustomer() Customer {
	var customer Customer
	customer.ID = 1
	customer.ShippingAddress = "123 Main St, New York, NY 10001"
	customer.BillingAddress = "123 Main St, New York, NY 10001"

	return customer
}
