package main

// IntegerHeap - a slice of int type
type IntegerHeap []int

// Attach methods to the IntegerHeap type

// Len - gets the length of integerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// Less - checks if element of i index is less than j
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// Swap - swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

// Push - pushes the item with the interface to the heap
func (iheap *IntegerHeap) Push(heapItem interface{}) {
	// we want to make an effect to the heap in memory and
	// not the copy which the function does, we make a reference
	// to the heap by getting a pointer from the item or data structure
	// then deference to assign value to the address where
	// the data is stored in memory
	// because heapItem is an empty interface, means that
	// any data can satisfy the interface since no method
	// is needed to be satisfied, then we extract the int value from it
	*iheap = append(*iheap, heapItem.(int))
}

// Pop - removes item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var (
		n, x1 int
	)

	// hold a reference to the 
	previous := *iheap
}