/*
Package main implements a priority queue data structure.

Queues are commonly used for storing tasks that need to be done
or incoming HTTP requests that need to be processed by a server.
Handling interruptions in real-time systems, call handling,
and CPU task scheduling are good examples for using queues.
*/
package main

import (
	"fmt"
)

// Queue data structure, it will be a list of Orders.
type Queue []*Order

// Order data type
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// newOrder function contructs a new order.
func newOrder(priority, quantity int, product, customerName string) *Order {
	return &Order{
		priority: priority,
		quantity: quantity,
		product: product,
		customerName: customerName,
	}
}

// Add method takes an order and adds it to the queue based on priority.
func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = Queue{order}
	} else {
		var appended bool

		for i, existingOrder := range *q {
			if order.priority > existingOrder.priority {
				// since order has a higher priority than the existing order
				// make a slice till the index
				start := (*q)[:i]

				// make anaother slice of the remaining queue starting from the
				// index we noticed the change in priority
				end := (*q)[i:]

				// set new order at the beginning of the remaining queue
				start = append(start, order)

				// concatenate or join all the queue into one with the
				*q = append(start, end...)
				appended = true
			}
		}

		if !appended {
			*q = append(*q, order)
		}
	}
}

func main() {
	q := make(Queue, 0)

	order1 := newOrder(2, 20, "Computer", "Andela")
	order2 := newOrder(1, 30, "Memory", "Westgate")

	q.Add(order1)
	q.Add(order2)

	for _, order := range q {
		fmt.Printf("Order %v \n", order)
	}
}
