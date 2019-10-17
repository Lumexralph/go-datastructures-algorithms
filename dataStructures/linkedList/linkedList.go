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

func main() {
	linkedList := LinkedList{}

	linkedList.AddToHead(1)
	fmt.Printf("The head node has a property %d\n", linkedList.headNode.property)

	linkedList.AddToHead(3)
	fmt.Printf("The head node has a property %d\n", linkedList.headNode.property)
}