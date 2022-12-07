package main

import (
	"container/list"
	"fmt"
)

func main() {
	q := newQueue()
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	q.enqueue(9)

	fmt.Println(q.dequeue())
	fmt.Printf("data: %+v\n", q.items())
}

type queue struct {
	list *list.List
}

func newQueue() queue {
	return queue{
		list: list.New(),
	}
}

func (q queue) enqueue(item int) {
	_ = q.list.PushBack(item)
}

func (q queue) dequeue() int {
	el := q.list.Front()
	if el == nil {
		return 0
	}
	val := q.list.Remove(el)
	return val.(int)
}

func (q queue) items() []int {
	var data []int
	el := q.list.Front()
	for ; el != nil; el = el.Next() {
		data = append(data, el.Value.(int))
	}
	return data
}

func (q queue) isEmpty() bool {
	return q.list.Front() == nil
}

//func main() {
//	input1 := [][]int{
//		{3, 2},
//		{3, 0},
//		{2, 0},
//		{2, 1},
//	}
//	fmt.Printf("subset: %v\n", topologicalSort(4, input1))
//
//	input2 := [][]int{
//		{4, 2},
//		{4, 3},
//		{2, 0},
//		{2, 1},
//		{3, 1},
//	}
//	fmt.Printf("subset: %v\n", topologicalSort(5, input2))
//
//	input3 := [][]int{
//		{6, 4},
//		{6, 2},
//		{5, 3},
//		{5, 4},
//		{3, 0},
//		{3, 1},
//		{3, 2},
//		{4, 1},
//	}
//	fmt.Printf("subset: %v\n", topologicalSort(7, input3))
//}

func topologicalSort(vertices int, edges [][]int) []int {
	// initialize the graph
	graph := make(map[int][]int)

	// initialize incomingEdges count
	incomingEdges := make(map[int]int)

	for i := 0; i < vertices; i++ {
		incomingEdges[i] = 0
		graph[i] = []int{}
	}

	// build the graph and incoming edges.
	for _, edge := range edges {
		vertex, child := edge[0], edge[1]
		graph[vertex] = append(graph[vertex], child)
		incomingEdges[child] += 1
	}

	// create a queue of graph sources.
	sources := newQueue()

	// build first set of sources
	for vertex, count := range incomingEdges {
		if count == 0 {
			sources.enqueue(vertex)
		}
	}

	var sortedOrder []int
	for !sources.isEmpty() {
		vertex := sources.dequeue()
		sortedOrder = append(sortedOrder, vertex)

		adjacencies := graph[vertex]
		for _, adjacent := range adjacencies {
			incomingEdges[adjacent] -= 1

			if incomingEdges[adjacent] == 0 {
				sources.enqueue(adjacent)
			}
		}
	}

	// if the sortedOrder is greater than the vertices, we have a cylce
	if len(sortedOrder) != vertices {
		return []int{} // or panic
	}

	return sortedOrder
}

// input: tasks = 4, prerequisite = [[3, 1], [1, 2], [3, 2]], output = true/false
// 3->2, 3->1->0, 3->2->1->0

func scheduleTask(numTask int, prerequisites [][]int) bool {
	// initialize graph
	graph := make(map[int][]int)
	incomingVertices := make(map[int]int) // vertex : count

	// build graph
	for _, pre := range prerequisites {
		parent, child := pre[0], pre[1]
		graph[parent] = append(graph[parent], child)

		if _, ok := incomingVertices[parent]; !ok {
			incomingVertices[parent] = 0
		}
		incomingVertices[child] += 1
	}

	// queue for sources
	sources := newQueue()

	for vertex, count := range incomingVertices {
		if count == 0 {
			sources.enqueue(vertex)
		}
	}

	var scheduleOrder []int
	for !sources.isEmpty() {
		task := sources.dequeue()
		scheduleOrder = append(scheduleOrder, task)

		for _, child := range graph[task] {
			incomingVertices[child] -= 1
			if incomingVertices[child] == 0 {
				sources.enqueue(child)
			}
		}
	}

	fmt.Println(scheduleOrder)
	return len(scheduleOrder) == numTask
}
