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

func main() {
	doublyLinkedList := DoublyLinkedList{}

	fmt.Printf("The node between properties %v", doublyLinkedList.NodeBetweenValues("a", "b"))
}