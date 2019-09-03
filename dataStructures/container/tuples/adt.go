package main

import (
	// "errors"
	"fmt"
)

// get the power series of integer a and returns
// tuple of square a and cube of a
func powerSeries(a int) (square int, cube int, err error) {

	square = a * a
	cube = a * a * a
	// err = errors.New("something went wrong")
	err = nil

	return
}

func main() {

	var (
		square int
		cube   int
		err    error
	)

	square, cube, err = powerSeries(5)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Square: %d - Cube: %d \n", square, cube)
}
