package v2

type stack struct {
	top *Node
}

func newStack(val int) *stack {
	return &stack{
		top: &Node{
			data: val,
		},
	}
}

func (s *stack) push(val int) {
	node := &Node{
		data: val,
	}

	if s.isEmpty() {
		s.top = node
	} else {
		node.next = s.top
		s.top = node
	}
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic("empty stack")
	}
	temp := s.top
	next := s.top.next
	s.top = next

	return temp.data
}

func (s *stack) peek() int {
	if s.isEmpty() {
		panic("empty stack")
	}
	return s.top.data
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func sortStack(s *stack) *stack {
	if s.isEmpty() {
		return &stack{}
	}

	currentElem := s.pop()
	sortedStack := new(stack)

	for !s.isEmpty() {
		if currentElem > s.peek() {
			// go further down the stack
			tempStack := new(stack)
			// move the peeked to the temp stack
			tempStack.push(s.pop())

			for !s.isEmpty() {
				if currentElem > s.peek() {
					tempStack.push(s.pop())
					continue
				} else {
					sortedStack.push(s.peek())
				}
			}

			sortedStack.push(currentElem)
			currentElem = tempStack.pop()
			s = tempStack
		}
	}
	sortedStack.push(currentElem)

	return s
}
