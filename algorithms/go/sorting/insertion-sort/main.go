// Package has the implementation of insertion sort
package main

import "fmt"

func main() {
	insertionSort([]int{5, 2, 4, 6, 1, 3})
}

// it assumes the input is a sequence of int
func insertionSort(input []int) []int {
	if len(input) == 1 {
		return input
	}

	// only the array or better said sub-array before
	// index j are sorted i.e towards the left and subarray
	// after the j index to the right are left but eventually get sorted
	for j := 0; j < len(input); j++ {
		// the key is used to compare with values towards
		// its left in the inner loop
		key := input[j]
		i := j - 1

		// if the values are bigger than key,
		// move them 1 position to the right
		for i >= 0 && input[i] > key {
			input[i+1] = input[i]
			i--
		}
		input[i+1] = key
	}

	fmt.Printf("Sorted=%v", input)
	return input
}
