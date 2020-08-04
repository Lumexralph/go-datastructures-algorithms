package linkedlist

import "fmt"

type nameNode struct {
	name string
	next *nameNode
}

func searchNode(n *nameNode, name string) *nameNode {
	// base case of when we are at the end
	if n == nil {
		return nil
	}
	if n.name == name {
		return n
	}

	// keep going
	return searchNode(n.next, name)
}

func insertList(n *nameList, name string) {
	newName := &nameNode{name, nil}
	if n.head == nil {
		n.head = newName
		return
	}
	previousHeader := n.head
	// link with the previous head and insert at the head
	newName.next = previousHeader
	n.head = newName

}

type nameList struct {
	head *nameNode
}

// newNameList creates a new name linked lists.
func newNameList(names ...string) *nameList {
	l := &nameList{}

	for _, name := range names {
		l.insertList(name)
	}

	return l
}

func (l *nameList) searchList(name string) *nameNode {
	return searchNode(l.head, name)
}

func (l *nameList) insertList(name string) {
	newName := &nameNode{name, nil}
	if l.head == nil {
		l.head = newName
		return
	}
	previousHeader := l.head
	// link with the previous head and insert at the head
	newName.next = previousHeader
	l.head = newName
}

func listNodes(n *nameNode) {
	if n == nil {
		return
	}
	fmt.Println("name: ", n.name)
	listNodes(n.next)
}

func (l *nameList) listEntries() {
	// start from the head of the node
	listNodes(l.head)
}

func nodesBeforeItem(n *nameNode, name string) *nameNode {
	if (n == nil) || n.next == nil {
		fmt.Println("no nodes before a nil list")
		return nil
	}

	if n.next.name == name { // look ahead
		return n
	}

	return nodesBeforeItem(n.next, name)
}

func (l *nameList) deleteList(name string) {
	// search if the node exists
	foundNode := l.searchList(name)

	if foundNode != nil {
		// get the nodes before it
		nodeBefore := nodesBeforeItem(l.head, name)
		if nodeBefore != nil {
			nodeBefore.next = foundNode.next
		} else { // it means the foundNode is the head node
			l.head = foundNode.next
		}
	}
}

// func main() {
// 	nl := &nameList{}
// 	insertList(nl, "olumide")
// 	insertList(nl, "femi")
// 	insertList(nl, "wale")
// 	insertList(nl, "adenike")

// 	fmt.Println(searchList(nl.head, "waley"))
// 	fmt.Println("print name entries: ")
// 	listEntries(nl.head)

// 	fmt.Println("remove a node: ")
// 	deleteList(nl, "adenike")

// 	fmt.Println("print updated entry: ")
// 	listEntries(nl.head)

// }
