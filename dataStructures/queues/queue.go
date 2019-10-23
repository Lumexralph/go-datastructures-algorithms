package main

// Queue data structure, it will be a list of Orders
type Queue []*Order

// Order data type
type Order struct {
	priority int
	quantity int
	product string
	customerName string
}

// New methos creates a new order with required properties
func (order *Order) New(priority, quantity int, product, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}