package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
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

			if !reflect.DeepEqual(got, tc.want) {
				// It's important to print the name of the function being tested _and_ the input
				// arguments too. Even if the test name is TestNewOrder, an error message
				// should contain all the information needed to debug it.
				t.Errorf("newOrder(%d, %d, %q, %q) got %+v; want %+v", tc.priority, tc.quantity, tc.product, tc.customerName, got, tc.want)
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

			if diff := cmp.Diff(gotQuantities, tc.wantQuantities); diff != "" {
				t.Errorf("Queue.Add() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
