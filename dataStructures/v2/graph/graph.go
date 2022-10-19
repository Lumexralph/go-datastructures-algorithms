package graph

import "fmt"

// Graph represents graph structure.
// This structure assumes the nodes have unique names.
type Graph struct {
	nodes map[string]*Node
}

func newGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}

func (g *Graph) insertNode(node *Node) {
	if _, ok := g.nodes[node.name]; ok {
		return
	}
	g.nodes[node.name] = node
}

func (g *Graph) insertAdjacent(node, adjacent string) {
	if _, ok := g.nodes[node]; !ok {
		panic("node does not exist")
	}

	adj := &Node{
		name:     adjacent,
		visited:  false,
		adjacent: nil,
	}
	if _, ok := g.nodes[adjacent]; !ok {
		g.insertNode(adj)
	}

	n := g.nodes[node]
	n.adjacent = append(n.adjacent, adj)
}

func (g *Graph) getAdjacent(node string) []*Node {
	if _, ok := g.nodes[node]; !ok {
		panic("node does not exist")
	}
	return g.nodes[node].adjacent
}

// Node represents individual node/vertex in a graph.
type Node struct {
	name     string
	visited  bool
	adjacent []*Node
}

var found bool

func depthFirstSearch(root *Node, key string) bool {
	if root == nil {
		return false
	}

	if root.name == key {
		found = true
		return found
	}
	root.visited = true

	for _, node := range root.adjacent {
		if node.name == key {
			fmt.Println("found: ", node.name)
			found = true
			return found
		}

		if !node.visited {
			node.visited = true
			depthFirstSearch(node, key)
		}
	}

	return found
}

func visit(node *Node) {
	var found bool
}
