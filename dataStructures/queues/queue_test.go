package main

import (
	"fmt"
	"testing"
)

func TestNewOrder(t *testing.T) {
	cases := []struct{
		name, product, customerName string
		priority, quantity int
		want *Order
	}{
		{"create new order", "macbook", "andela", 1, 20, &Order{1, 20, "macbook", "andela"}},
		{"create another order", "windows", "apple", 2, 30, &Order{2, 30, "windows", "apple"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ord := newOrder(tc.priority, tc.quantity, tc.product, tc.customerName)
			if ord.customerName != tc.want.customerName {
				t.Errorf("got %s, want %s", ord.customerName, tc.want.customerName)
			}

			if ord.product != tc.want.product {
				t.Errorf("got %s, want %s", ord.product, tc.want.product)
			}

			if ord.priority != tc.want.priority {
				t.Errorf("got %d, want %d", ord.priority, tc.want.priority)
			}

			if ord.quantity != tc.want.quantity {
				t.Errorf("got %d, want %d", ord.quantity, tc.want.quantity)
			}
		})
	}
}

func TestAddToQueue(t *testing.T) {
	q := make(Queue, 0)
	cases := []struct{
		o1, o2 *Order
		want int
	}{
		{&Order{1, 20, "macbook", "andela"}, &Order{2, 30, "windows", "apple"}, 2},
		{&Order{3, 20, "macbook", "andela"}, &Order{4, 30, "windows", "apple"}, 4},
		{&Order{5, 20, "macbook", "andela"}, &Order{6, 30, "windows", "apple"}, 6},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("orders with priority %d and %d", tc.o1.priority, tc.o2.priority), func(t *testing.T) {
			q.Add(tc.o1)
			q.Add(tc.o2)

			if q[0].priority != tc.want {
				t.Errorf("got %d, want %d", q[0].priority, tc.want)
			}
		})
	}

}