// Package binarytree is holds implementation for Binary search tree
// to store the words of dictionary.
// The data structure by default will sort entries, it has
// fast search and updates.
//
// The Binary Search Tree will support 4 operations:
// 1. Searching 2. Traversal 3. Insertion 4. Deletion
package binarytree

//import "fmt"

// Tree is the basic structure or node in a binary search tree.
type Tree struct {
	parent *Tree  // parent can be null i.e Root tree
	left   *Tree  // left child or left subtree
	right  *Tree  // right child or right subtree
	Item   string // data held by the node
}

// Search searches for an item in the tree.
func (t *Tree) Search(tree *Tree, item string) *Tree {
	if tree == nil {
		return nil
	}

	if item == tree.Item { // found the item
		return tree
	} else if item < tree.Item { // continue search on the left subtree
		return t.Search(tree.left, item)
	} else { // it means item is on the right side
		return t.Search(tree.right, item)
	}
}

// BinarySearchTree to store the words of dictionary.
type BinarySearchTree struct {
	Root *Tree
}

// NewBinaryTree creates a new binary search tree.
func NewBinaryTree(word string) *BinarySearchTree {
	RootTree := &Tree{
		parent: nil,
		left:   nil,
		right:  nil,
		Item:   word,
	}
	return &BinarySearchTree{Root: RootTree}
}

// Search searches for an item, it expects to start from the Root.
func (bt *BinarySearchTree) Search(item string) string {

	 if result := bt.Root.Search(bt.Root, item); result != nil { // found!
	 	return result.Item
	 }
	 return ""
}

//FindMinimum returns the smallest or lowest item item in the tree.
//
// it works by moving continually through the left subtree till
// we reach the last tree which we can in a way call the leaf.
func (bt *BinarySearchTree) FindMinimum() *Tree {
	// start from the Root and move to the left subtrees
	minTree := bt.Root

	for minTree.left != nil {
		minTree = minTree.left
	}
	return minTree
}

// FindMaximum returns the largest or highest item item in the tree.
func (bt *BinarySearchTree) FindMaximum() *Tree {
	// start from the Root and move right wards
	maxTree := bt.Root

	for maxTree.right != nil {
		maxTree = maxTree.right
	}
	return maxTree
}

// InOrderTraversal traverse in the order [smallest ... largest] [allLeft ... allRight]
func (bt *BinarySearchTree) InOrderTraversal(t *Tree, processItem func (string)) {
	if t != nil {
		bt.InOrderTraversal(t.left, processItem)
		// process the item in the tree
		processItem(t.Item)
		bt.InOrderTraversal(t.right, processItem)
	}
}

func (bt *BinarySearchTree) PreOrderTraversal(t *Tree, processItem func (string)) {
	if t != nil {
		// process the item in the tree
		processItem(t.Item)
		bt.PreOrderTraversal(t.left, processItem)
		bt.PreOrderTraversal(t.right, processItem)
	}
}

// PostOrderTraversal traverses in the order [lastInserted, beforeLastInserted, ..., secondInserted, firstInserted]
// It reverses the order of insertion.
func (bt *BinarySearchTree) PostOrderTraversal(t *Tree, processItem func (string)) {
	if t != nil {
		bt.PostOrderTraversal(t.left, processItem)
		bt.PostOrderTraversal(t.right, processItem)
		// process the item in the tree
		processItem(t.Item)
	}
}

func ProcessItem(itemCollection *[]string) func(string) {
	return func(item string) {
		*itemCollection = append(*itemCollection, item)
	}
}


func (bt *BinarySearchTree) RecursiveInsert(currentTree **Tree, item string, parent *Tree) {
	// currentTree is a pointer to the memory location(also a pointer) where the current node tree
	// is stored. This is needed because, when we get to a tree that is empty and we need to insert
	// the new item, we'll need to replace it with the new tree, this warrants having access to the location.

	// when we have come to the end of the search.
	if *currentTree == nil {
		newTree := &Tree{
			Item:   item,
			left:   nil,
			right:  nil,
			parent: parent,
		}
		// Attach to its parent, by getting the memory location where the currentTree.
		*currentTree = newTree
		return
	}

	// Recursively search for where to put the item.
	if item < (*currentTree).Item { // diverge to the left subtree
		bt.RecursiveInsert(&(*currentTree).left, item, *currentTree)
	} else {
		bt.RecursiveInsert(&(*currentTree).right, item, *currentTree)
	}
}

func (bt *BinarySearchTree) NormalInsert(item string) {
	if bt.Root == nil {
		bt.Root = &Tree{
			Item: item,
		}
		return
	}
	currentTree := bt.Root

insertionLoop:
	for {
		switch {
		// go to the left subtree
		case item < currentTree.Item:
			if currentTree.left == nil { // at the end of the left subtrees
				// attach the new node
				currentTree.left = &Tree{
					Item:   item,
					parent: currentTree,
				}
				break insertionLoop
			} else {
				currentTree = currentTree.left
			}
			break
		// go to the right subtree
		case item > currentTree.Item:
			if currentTree.right == nil { // at the end of the right subtree
				// attach the new node
				currentTree.right = &Tree{
					Item:   item,
					parent: currentTree,
				}
				break insertionLoop
			} else {
				currentTree = currentTree.right
			}
			break
		default:
			break insertionLoop
		}
	}
}

// produceTreeItems uses normal insertion for items.
func ProduceTreeItems(items ...string) *BinarySearchTree {
	if len(items) == 0 {
		return NewBinaryTree("")
	}

	bst := NewBinaryTree(items[0])
	for _, item := range items[1:] {
		bst.NormalInsert(item)
	}
	return bst
}

func ProduceTreeItemsRecursive(items ...string) *BinarySearchTree {
	if len(items) == 0 {
		return NewBinaryTree("")
	}

	bst := NewBinaryTree(items[0])
	for _, item := range items[1:] {
		bst.RecursiveInsert(&bst.Root, item, nil)
	}
	return bst
}

//func main() {
//	// Create a new binary search tree
//	bst := NewBinaryTree("game")
//	bst.RecursiveInsert(&bst.Root, "fame", nil)
//	bst.NormalInsert("boy")
//
//	fmt.Println("search: ", bst.Search("boy"))
//
//	//fmt.Println("inorder traversal: ")
//	//bst.InOrderTraversal(bst.Root, processItem)
//
//	//fmt.Println("preorder traversal: ")
//	//itemCollection := new([]string)
//	//bst.PreOrderTraversal(bst.Root, processItem(itemCollection))
//	////
//	////fmt.Println("postorder traversal: ")
//	////bst.PostOrderTraversal(bst.Root, processItem)
//	//
//	//maxNode := bst.FindMaximum()
//	//minNode := bst.FindMinimum()
//	//
//	//fmt.Println("maximum: ", maxNode.item)
//	//fmt.Println("minimum: ", minNode.item)
//	//fmt.Printf("itemcollection:  %q\n", *itemCollection)
//}
