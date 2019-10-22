package main

// Set data structure
type Set struct {
	integerMap map[int]bool
}

// New method to create a new map in the set
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}
