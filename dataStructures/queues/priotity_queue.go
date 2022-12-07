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

// insertionSlot method will help to pick a slot to fix the next order
// with complexity of Olog(n)
func insertionSlot(q Queue, value int) int {
	low := 0
	high := len(q) - 1

	for low <= high {
		mid := (high + low) / 2

		fmt.Println(low, high)

		switch {
		case q[mid].priority == value:
			if i := (mid + 1); i < len(q) && q[i].priority == value {
				low = i
			} else {
				return i
			}
		case q[mid].priority < value:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	return low

}

// Add method takes an order and adds it to the queue based on priority.
func (q *Queue) Add(order *Order) {
	insertPoint := insertionSlot(*q, order.priority)

	if insertPoint > len(*q)-1 {
		// when the insert point is at the end of the queue
		*q = append(*q, order)
	} else {
		start := copyQ((*q)[:insertPoint])
		start = append(start, order)
		end := copyQ((*q)[insertPoint:])

		// recreate the whole queue
		*q = append(start, end...)
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

// func main() {
//	h := &maxHeap{2, 1, 5}
//	heap.Init(h)
//	heap.Push(h, 3)
//
//	fmt.Println("maxHeap.Peek: ", h.Peek())
//
//	h1 := &minHeap{2, 1, 5}
//	heap.Init(h1)
//	heap.Push(h1, 3)
//
//	fmt.Println("minHeap.Peek: ", h1.Peek())
// }

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h maxHeap) Peek() int {
	return h[0]
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minHeap) Peek() int {
	return h[0]
}
