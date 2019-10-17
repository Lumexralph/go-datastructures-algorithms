package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

// AddToHead method adds the node to the start of the linked list
func(linkedList *LinkedList) AddToHead(property int) {
	node := Node{
		property: property,
		nextNode: nil,
	}
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

// IterateList method iterates over a LinkedList
func (linkedList *LinkedList) IterateList() {
	node := linkedList.headNode
	for node != nil {
		fmt.Printf("%d ", node.property)
		node = node.nextNode
	}
	fmt.Println("")
}

// GetLastNode method returns the last node in the linkedlist
func (linkedList *LinkedList) GetLastNode() *Node {
	var lastNode *Node

	node := linkedList.headNode

	for node != nil {
		lastNode = node
		node = node.nextNode
	}

	return lastNode
}

// AddToEnd method adds the node at the end of the list
func(linkedList *LinkedList) AddToEnd(property int) {
	// create new node as the last node
	newNode := &Node{
		property: property,
		nextNode: nil,
	}

	// get the last node
	lastNode := linkedList.GetLastNode()
	if lastNode != nil {
		lastNode.nextNode = newNode
	}
}

// GetNodeWithValue method of LinkedList returns the
// node with the property value
func (linkedList *LinkedList) GetNodeWithValue(property int) *Node {
	var foundNode *Node

	node := linkedList.headNode

	for node != nil {
		if node.property == property {
			foundNode = node
			break
		}
		node = node.nextNode
	}

	return foundNode
}

// AddAfter method adds the node after a specific node
func (linkedList *LinkedList) AddAfter(nodeProperty int, newProperty int) {
	// create a new node
	newNode := &Node{
		property: newProperty,
		nextNode: nil,
	}
	// attach the nextNode of the current node to the new node
	currentNode := linkedList.GetNodeWithValue(nodeProperty)

	// attach the nextNode of the new node to the initial nextNode
	if currentNode != nil {
		newNode.nextNode = currentNode.nextNode
		currentNode.nextNode = newNode
	}
}

func main() {
	linkedList := LinkedList{}

	linkedList.AddToHead(1)
	fmt.Printf("The head node has a property %d\n", linkedList.headNode.property)

	linkedList.AddToHead(3)
	fmt.Printf("The head node has a property %d\n", linkedList.headNode.property)

	linkedList.AddToEnd(5)
	linkedList.AddToEnd(7)
	fmt.Printf("Found Node %v \n", linkedList.GetNodeWithValue(7).property)
	linkedList.AddAfter(5, 9)
	linkedList.IterateList()
}