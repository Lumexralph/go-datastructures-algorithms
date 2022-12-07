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

// InterSect method returns another set with elements that exists in two sets
func (set *Set) InterSect(anotherSet *Set) *Set {
	intersectSet := &Set{}

	intersectSet.New()

	for element := range set.integerMap {
		if anotherSet.ContainsElement(element) {
			intersectSet.AddElement(element)
		}
	}

	return intersectSet
}

// Union method returns a set that consists of elements of two sets
func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()

	for element := range set.integerMap {
		unionSet.AddElement(element)
	}

	for element := range anotherSet.integerMap {
		unionSet.AddElement(element)
	}

	return unionSet
}

func main() {
	set := &Set{}
	anotherSet := &Set{}

	set.New()
	set.AddElement(2)
	set.AddElement(3)
	set.AddElement(5)
	set.AddElement(5)

	anotherSet.New()
	anotherSet.AddElement(4)
	anotherSet.AddElement(6)
	anotherSet.AddElement(2)
	anotherSet.AddElement(3)

	fmt.Println(set.ContainsElement(2))
	fmt.Println(set.ContainsElement(4))
	fmt.Println("Union ", set.Union(anotherSet))
	fmt.Println("Intersect ", set.InterSect(anotherSet))

	input := []int{1, 5, 3}
	fmt.Printf("subset: %v\n", createListSubset(input))
}

func createListSubset(sets []int) [][]int {
	var listSubset [][]int

	listSubset = append(listSubset, []int{})

	for _, val := range sets {
		for _, subset := range listSubset {
			var newSet []int
			newSet = append(newSet, subset...)
			newSet = append(newSet, val)
			listSubset = append(listSubset, newSet)
		}
	}

	return listSubset
}
