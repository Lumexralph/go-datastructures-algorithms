package main

import (
	"fmt"
)

func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}

func main() {
	slice := []int{1, 3, 4, 5}
	slice = append(slice, 3)

	fmt.Println("Capacity ", cap(slice))
	fmt.Println("Length ", len(slice))

	twiceValue(slice)

	for i, value := range slice {
		fmt.Printf("New index with value %d - %d \n", i, value)
	}
}