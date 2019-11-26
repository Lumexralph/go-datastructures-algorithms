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
		priority:     priority,
		quantity:     quantity,
		product:      product,
		customerName: customerName,
	}
}

// create a copy of the slice to avoid referencing same memory location
// when adding new order to the slice
func copyQ(q Queue) Queue {
	return append(Queue{}, q...)
}

// Add method takes an order and adds it to the queue based on priority.
func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = Queue{order}
	} else {
		var appended bool

		for i, existingOrder := range *q {
			if order.priority < existingOrder.priority {
				// since order has a higher priority than the existing order
				// make a slice till the index
				start := (*q)[:i]

				// make anaother slice of the remaining queue starting from the
				// index we noticed the change in priority
				end := copyQ((*q)[i:])

				// set new order at the beginning of the remaining queue
				start = append(start, order)

				// concatenate or join all the queue into one with the
				*q = append(start, end...)
				appended = true
				break
			}

			if order.priority == existingOrder.priority {
				nextIndex := i + 1

				// peek to check next order in the queue
				if nextIndex < len(*q) && order.priority == (*q)[nextIndex].priority {
					// move on in the queue to ensure no order has same priority
					continue
				}

				// if the priority levels are same e.g 2 == 2
				// append the order after the existing order with same priority
				// make a slice including the existing order
				start := (*q)[:nextIndex]

				// make another slice to have the remaining orders
				end := copyQ((*q)[nextIndex:])

				// add the new order
				start = append(start, order)

				// concatenate or join all the queue into one with the
				*q = append(start, end...)
				appended = true
				break

			}

		}

		if !appended {
			*q = append(*q, order)
		}
	}
}

func main() {
	q := new(Queue)
	// 3, 8, 2, 5, 3, 3, 2, 3, 2
	order1 := newOrder(2, 20, "Computer", "Andela")
	order2 := newOrder(1, 30, "Memory", "Westgate")
	order3 := newOrder(1, 30, "cpu", "Westside")
	order4 := newOrder(2, 30, "disk", "southside")
	order5 := newOrder(3, 30, "game", "easthside")
	order6 := newOrder(8, 30, "time", "northside")
	order7 := newOrder(2, 30, "nadia", "southside")
	order8 := newOrder(2, 30, "gamers", "southside")
	order9 := newOrder(3, 30, "psygame", "easthside")

	q.Add(order1)
	q.Add(order2)
	q.Add(order3)
	q.Add(order4)
	q.Add(order5)
	q.Add(order6)
	q.Add(order7)
	q.Add(order8)
	q.Add(order9)

	for _, order := range *q {
		fmt.Printf("Order %v \n", order)
	}
}
