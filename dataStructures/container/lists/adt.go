package main

import (
	"container/list"
	"fmt"
)

func main() {

	// a doubly-linked list is created
	var intList list.List

	// add elements to the list
	intList.PushBack(24)
	intList.PushBack(50)
	intList.PushBack(40)
	intList.PushBack(4)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
