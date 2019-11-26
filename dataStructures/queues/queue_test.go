package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewOrder(t *testing.T) {
	cases := []struct {
		name, product, customerName string
		priority, quantity          int
		want                        *Order
	}{
		{"create new order", "macbook", "andela", 1, 20, &Order{1, 20, "macbook", "andela"}},
		{"create another order", "windows", "apple", 2, 30, &Order{2, 30, "windows", "apple"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// In tests it's common to call the value that you're testing "got"
			// If there are lots of them then "gotOrder" "gotQueue" etc. but in this test it doesn't matter
			// We do it because then someone reading the code immediately knows that this is the value being tested
			got := newOrder(tc.priority, tc.quantity, tc.product, tc.customerName)
			if got.customerName != tc.want.customerName {
				t.Errorf("newOrder(%d, %d, %q, %q) got +%s; want +%s", tc.priority, tc.quantity, tc.product, tc.customerName, got.customerName, tc.want.customerName)
			}

			if got.product != tc.want.product {
				t.Errorf("newOrder(%d, %d, %q, %q) got +%s; want +%s", tc.priority, tc.quantity, tc.product, tc.customerName, got.product, tc.want.product)
			}

			if got.priority != tc.want.priority {
				t.Errorf("newOrder(%d, %d, %q, %q) got +%v; want +%v", tc.priority, tc.quantity, tc.product, tc.customerName, got.priority, tc.want.priority)
			}

			if got.quantity != tc.want.quantity {
				t.Errorf("newOrder(%d, %d, %q, %q) got +%v; want +%v", tc.priority, tc.quantity, tc.product, tc.customerName, got.quantity, tc.want.quantity)
			}

			if !reflect.DeepEqual(got, tc.want) {
				// It's important to print the name of the function being tested _and_ the input
				// arguments too. Even if the test name is TestNewOrder, an error message
				// should contain all the information needed to debug it.
				t.Errorf("newOrder(%d, %d, %q, %q) got %+v; want %+v", tc.priority, tc.quantity, tc.product, tc.customerName, got, tc.want)
			}
		})
	}
}

func TestAddToQueue(t *testing.T) {
	cases := []struct {
		o1, o2 *Order
		want   int
	}{
		{&Order{1, 20, "macbook", "andela"}, &Order{2, 30, "windows", "apple"}, 1},
		{&Order{3, 20, "macbook", "andela"}, &Order{4, 30, "windows", "apple"}, 3},
		{&Order{5, 20, "macbook", "andela"}, &Order{6, 30, "windows", "apple"}, 5},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("orders with priority %d and %d", tc.o1.priority, tc.o2.priority), func(t *testing.T) {
			q := new(Queue)
			q.Add(tc.o1)
			q.Add(tc.o2)

			if (*q)[0].priority != tc.want {
				t.Errorf("got %d, want %d", (*q)[0].priority, tc.want)
			}
		})
	}

}

func TestAddToEmptyQueue(t *testing.T) {
	cases := []struct {
		name           string
		priorities     []int // order quantities will be equal to their index in the slice
		wantQuantities []int
	}{
		{
			name:           "multiple equal priorities",
			priorities:     []int{3, 8, 2, 5, 3, 3, 2, 3, 2},
			wantQuantities: []int{2, 6, 8, 0, 4, 5, 7, 3, 1}, //It's an argsort https://stackoverflow.com/questions/17901218/numpy-argsort-what-is-it-doing
			// (if the slice was sorted, where the element will be using their index in the priorities slice)
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			q := &Queue{}

			var orders []*Order
			for i, p := range tc.priorities {
				o := newOrder(p, i, "", "") // using the index of the slice as the quantity
				orders = append(orders, o)
				q.Add(o)
			}

			var gotQuantities []int
			for _, o := range *q {
				gotQuantities = append(gotQuantities, o.quantity)
			}

			if !reflect.DeepEqual(gotQuantities, tc.wantQuantities) {
				t.Errorf("after adding Orders %+v to empty Queue; got Queue quantities %d; want %d", orders, gotQuantities, tc.wantQuantities)
			}
		})
	}
}
