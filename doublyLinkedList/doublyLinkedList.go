package main

import (
	"fmt"
)

// Node for a doubly linked list
type Node struct {
	property string
	prevNode *Node
	nextNode *Node
}

// DoublyLinkedList container for the nodes
type DoublyLinkedList struct {
	headNode *Node
}

// NodeBetweenValues method returns the node that has a property
// lying between the firstProperty and secondProperty values
func (doublyLinkedList *DoublyLinkedList) NodeBetweenValues(firstProperty string, secondProperty string) *Node {
	node := doublyLinkedList.headNode

	for node != nil {
		// check if the node is between the two properties
		if node.prevNode.property == firstProperty && node.nextNode.property == secondProperty {
			return node
		}

		node = node.nextNode
	}

	return nil
}

// AddToHead method of the doubly LinkedList sets the previousNode
// property of the current headNode of the linked list to the node
// that's added with property
func (doublyLinkedList *DoublyLinkedList) AddToHead(property string) {
	newNode := &Node{
		property: property,
		prevNode: nil,
		nextNode: nil,
	}
	headNode := doublyLinkedList.headNode

	if headNode != nil {
		newNode.nextNode = headNode
		headNode.prevNode = newNode
	}

	doublyLinkedList.headNode = newNode
}

func main() {
	doublyLinkedList := DoublyLinkedList{}
	doublyLinkedList.AddToHead("c")

	fmt.Printf("The node between properties %v", doublyLinkedList.NodeBetweenValues("a", "b"))
}