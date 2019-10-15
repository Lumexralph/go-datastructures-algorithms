package main

import (
	"fmt"
)

var arr = [5]int{1, 2, 3, 4, 5}

func main() {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Printing Elements %d \n", arr[i])
	}

	// use range to access index and value
	for i, value := range arr {
		fmt.Printf("Index %d - value %d using range \n", i, value)
	}
}
