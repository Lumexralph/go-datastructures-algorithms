package main

import (
	"fmt"
)

func main() {
	rows := 10
	cols := 12

	twoDSlices := make([][]int, rows)
	for i :=range twoDSlices {
		twoDSlices[i] = make([]int, cols)
	}

	fmt.Println(twoDSlices)
}