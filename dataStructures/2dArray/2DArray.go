package main

import (
	"fmt"
)

func main() {
	var twoDArray [8][8]int

	twoDArray[3][7] = 20
	twoDArray[2][7] = 30

	fmt.Println(twoDArray)
}