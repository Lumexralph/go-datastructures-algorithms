package main

import (
	"fmt"
)

// IComposite Interface
type IComposite interface {
	perform() string
}

// Leaflet struct (model, class)
type Leaflet struct {
	name string
}

// Leaflet perform method
func (leaf *Leaflet) perform() string {
	return "Leaflet :" + leaf.name
}

// Branch struct, it can have a leaf or contain more branches
type Branch struct {
	leafs    []*Leaflet
	name     string
	branches []*Branch
}

// Branch struct has the perform method that satisfies the IComposite interface
func (branch *Branch) perform() string {

	// go through the leafs to call the method on each leaf
	for _, leaf := range branch.leafs {
		fmt.Println(leaf.perform())
	}

	// go through the branches recursively to get to the node or leaf
	for _, branch := range branch.branches {
		fmt.Println(branch.perform())
	}

	return "Branch : " + branch.name
}

// Branch method to add a leaflet
func(branch *Branch) add(leaf *Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// Branch method to add new branch
func(branch *Branch) addBranch(newBranch *Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// method to get leaflets
func(branch *Branch) getLeaflets() []*Leaflet {
	return branch.leafs
}

func main() {
	var branch1 = &Branch{name: "Branch 1"}
	var branch2 = &Branch{name: "Branch 2"}
	var leaf1 = &Leaflet{name: "Leaflet 1"}
	var leaf2 = &Leaflet{name: "Leaflet 2"}

	branch1.add(leaf1)
	branch2.add(leaf2)
	branch1.addBranch(branch2)

	// perform the operation on the whole tree
	totalResult := branch1.perform()
	fmt.Println(totalResult)
}