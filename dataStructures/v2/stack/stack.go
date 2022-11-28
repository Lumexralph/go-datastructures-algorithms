package stack

import (
	"container/list"
	"fmt"
)

func main() {
	s := newStack()
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)
	s.push(9)

	fmt.Println(s.pop())
	fmt.Printf("data: %+v\n", s.items())
}

type stack struct {
	list *list.List
}

func newStack() stack {
	return stack{
		list: list.New(),
	}
}

func (s stack) push(item int) {
	_ = s.list.PushFront(item)
}

func (s stack) pop() int {
	el := s.list.Front()
	if el == nil {
		return 0
	}
	val := s.list.Remove(el)
	return val.(int)
}

func (s stack) items() []int {
	var data []int
	el := s.list.Front()
	for ; el != nil; el = el.Next() {
		data = append(data, el.Value.(int))
	}
	return data
}
