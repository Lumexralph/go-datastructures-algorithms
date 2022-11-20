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
