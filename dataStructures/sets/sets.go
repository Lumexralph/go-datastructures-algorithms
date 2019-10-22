package main

import (
	"fmt"
)

// Set data structure
type Set struct {
	integerMap map[int]bool
}

// New method to create a new map in the set
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// AddElement method adds the element to a set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// DeleteElement method deletes the element from integerMap using the delete method
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// ContainsElement method of the Set checks whether or not the element exists
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}

func main() {
	set := &Set{}

	set.New()
	set.AddElement(2)
	set.AddElement(3)
	set.AddElement(5)
	set.AddElement(5)

	fmt.Println(set.ContainsElement(2))
	fmt.Println(set.ContainsElement(4))

	fmt.Println(set)
}
