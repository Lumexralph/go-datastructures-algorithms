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

// copyQ creates a copy of the slice to avoid referencing same memory location
// when adding new order to the slice.
func copyQ(q Queue) Queue {
	return append(Queue{}, q...)
}

// Add method takes an order and adds it to the queue based on priority.
func (q *Queue) Add(order *Order) {
	var appended bool

	if len(*q) == 1 {
		if (*q)[0].priority > order.priority {
			*q = []*Order{order, (*q)[0]}
			appended = true
		}
	} else {
		// using binary-search to know index to insert the order.
		lowIndex := 0
		highIndex := len(*q) - 1
	QLoop:
		for lowIndex <= highIndex {

			midPoint := (lowIndex + highIndex) / 2
			guessOrder := (*q)[midPoint]

			switch {
			case order.priority < guessOrder.priority:
				// we move closer to the left-side
				highIndex = midPoint - 1
			case order.priority > guessOrder.priority:
				// we move to the right-side of the queue
				lowIndex = midPoint + 1

			case order.priority == guessOrder.priority || lowIndex == highIndex:
				insert := midPoint
				for insert < len(*q) && (*q)[insert].priority == order.priority {
					insert++
				}
				// make a slice till the index.
				start := (*q)[:insert]

				// make another slice to have the remaining orders.
				end := copyQ((*q)[insert:])

				// add the new order
				start = append(start, order)

				// concatenate or join all the queue into one with the.
				*q = append(start, end...)
				appended = true
				break QLoop
			}

		}

		if !appended {
			*q = append(*q, order)
		}
	}
}

func main() {
	q := new(Queue)

	order1 := newOrder(2, 20, "Computer", "Andela")
	order2 := newOrder(1, 30, "Memory", "Westgate")
	order3 := newOrder(1, 30, "cpu", "Westside")
	order4 := newOrder(2, 30, "disk", "southside")
	order5 := newOrder(3, 30, "game", "easthside")
	order6 := newOrder(8, 30, "time", "northside")
	order7 := newOrder(2, 30, "nadia", "southside")
	order8 := newOrder(2, 30, "gamers", "southside")
	order9 := newOrder(3, 30, "psygame", "easthside")
	order10 := newOrder(2, 37, "pc", "mixedreality")

	q.Add(order1)
	q.Add(order2)
	q.Add(order3)
	q.Add(order4)
	q.Add(order5)
	q.Add(order6)
	q.Add(order7)
	q.Add(order8)
	q.Add(order9)
	q.Add(order10)

	for _, order := range *q {
		fmt.Printf("Order %v \n", order)
	}
}
