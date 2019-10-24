package main

import (
	"fmt"
)

// Queue data structure, it will be a list of Orders
type Queue []*Order

// Order data type
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// New method creates a new order with required properties
func (order *Order) New(priority, quantity int, product, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

// Add method takes an order and adds it to the queue based on priority
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false

		for i, existingOrder := range *queue {
			if order.priority > existingOrder.priority {
				// since order has a higher priority than the existing order
				// make a slice till the index
				initialQueue := (*queue)[:i]

				// make anaother slice of the remaining queue starting from the
				// index we noticed the change in priority
				remainingQueue := (*queue)[i:]

				// create a new queue with the new order
				// set new order at the beginning of the remaining queue
				newQueue := Queue{order}
				newQueue = append(newQueue, remainingQueue...)

				// concatenate or join all the queue into one with the
				*queue = append(initialQueue, newQueue...)
				appended = true
			}
		}

		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	queue := make(Queue, 0)
	order1 := &Order{}
	order1.New(2, 20, "Computer", "Andela")

	order2 := &Order{}
	order2.New(1, 30, "Memory", "Westgate")

	queue.Add(order1)
	queue.Add(order2)

	for _, order := range queue {
		fmt.Printf("Order %v \n", order)
	}
}
