package v2

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func New(val int) *LinkedList {
	return &LinkedList{
		head: &Node{
			data: val,
			next: nil,
		},
	}
}

func (ll *LinkedList) insert(val int) {
	n := ll.head

	for n.next != nil {
		n = n.next
	}
	n.next = &Node{
		data: val,
		next: nil,
	}
}

func removeDuplicates(list *LinkedList) *LinkedList {
	var prevNode *Node = nil
	currNode := list.head
	record := make(map[int]struct{})

	for currNode.next != nil {
		if _, ok := record[currNode.data]; ok {
			prevNode.next = currNode.next
		}
		record[currNode.data] = struct{}{}
		prevNode = currNode
		currNode = currNode.next
	}

	return list
}

// Kth position to last element.
func kthToLast(llist *LinkedList, position int) int {
	if llist == nil {
		return -1
	}

	slowPointer := llist.head
	fastPointer := llist.head

	for slowPointer.next != nil {
		distance := 0
		for fastPointer != nil {
			if distance < position {
				fastPointer = fastPointer.next
				distance += 1
				continue
			}

			fastPointer = fastPointer.next
			if slowPointer.next != nil {
				slowPointer = slowPointer.next
			}
		}
		// anything that gets here means we are done.
		break
	}
	return slowPointer.data
}

func kthToLastRecursion(head *Node, position int) int {
	if head == nil {
		return 0
	}

	index := kthToLastRecursion(head.next, position) + 1
	if index == position {
		return head.data
	}

	return index
}

func partitionList(node *Node, val int) *Node {
	var beforeStart *Node = nil
	var beforeEnd *Node = nil

	var afterStart *Node = nil
	var afterEnd *Node = nil

	for node != nil {
		next := node.next
		node.next = nil
		if node.data < val {
			if beforeStart == nil {
				beforeStart = node
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = node
				beforeEnd = node
			}
		} else {
			if afterStart == nil {
				afterStart = node
				afterEnd = afterStart
			} else {
				afterEnd.next = node
				afterEnd = node
			}
		}
		node = next
	}

	if beforeStart == nil {
		return afterStart
	}

	beforeEnd.next = afterStart

	return beforeStart
}

func listSum(left, right *LinkedList) int {
	var leftOperand *Node = nil
	var rightOperand *Node = nil
	sum := 0

	if left != nil {
		leftOperand = left.head
	}
	if right != nil {
		rightOperand = right.head
	}

	for rightOperand != nil || leftOperand != nil {
		result := 0

		if rightOperand != nil && leftOperand != nil {
			result = rightOperand.data + leftOperand.data
		} else if rightOperand != nil {
			result = rightOperand.data
		} else {
			result = leftOperand.data
		}

		sum = (sum * 10) + result

		if rightOperand != nil {
			rightOperand = rightOperand.next
		}
		if leftOperand != nil {
			leftOperand = leftOperand.next
		}
	}

	return sum
}

func listSumReverseDigit(left, right *LinkedList) int {
	var leftOperand *Node = nil
	var rightOperand *Node = nil

	var rightDigit strings.Builder
	var leftDigit strings.Builder

	if left != nil {
		leftOperand = left.head
	}
	if right != nil {
		rightOperand = right.head
	}

	for leftOperand != nil || rightOperand != nil {
		if leftOperand != nil {
			str := fmt.Sprintf("%d%s", leftOperand.data, leftDigit.String())
			leftDigit.Reset()
			leftDigit.WriteString(str)
			leftOperand = leftOperand.next
		}
		if rightOperand != nil {
			str := fmt.Sprintf("%d%s", rightOperand.data, rightDigit.String())
			rightDigit.Reset()
			rightDigit.WriteString(str)
			rightOperand = rightOperand.next
		}
	}

	r, err := strconv.Atoi(rightDigit.String())
	fmt.Println(rightDigit.String(), leftDigit.String())
	if err != nil {
		return 0
	}

	l, err := strconv.Atoi(leftDigit.String())
	if err != nil {
		return 0
	}

	return r + l
}

func listSumDigit(left, right *LinkedList) int {
	var leftOperand *Node = nil
	var rightOperand *Node = nil

	var rightDigit strings.Builder
	var leftDigit strings.Builder

	if left != nil {
		leftOperand = left.head
	}
	if right != nil {
		rightOperand = right.head
	}

	for leftOperand != nil || rightOperand != nil {
		if leftOperand != nil {
			str := fmt.Sprintf("%s%d", leftDigit.String(), leftOperand.data)
			leftDigit.Reset()
			leftDigit.WriteString(str)
			leftOperand = leftOperand.next
		}
		if rightOperand != nil {
			str := fmt.Sprintf("%s%d", rightDigit.String(), rightOperand.data)
			rightDigit.Reset()
			rightDigit.WriteString(str)
			rightOperand = rightOperand.next
		}
	}

	r, err := strconv.Atoi(rightDigit.String())
	fmt.Println(rightDigit.String(), leftDigit.String())
	if err != nil {
		return 0
	}

	l, err := strconv.Atoi(leftDigit.String())
	if err != nil {
		return 0
	}

	return r + l
}

func isLoop(llist *LinkedList) bool {
	cycleLength := 0
	isCycle := false
	var cycleStart *Node = nil

	if llist == nil || llist.head == nil || llist.head.next == nil {
		return false
	}

	slow := llist.head
	fast := llist.head

	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if isCycle {
			cycleLength += 1
			if slow == cycleStart {
				fmt.Println("cycle length: ", cycleLength)
				return isCycle
			}
		}

		if fast == slow && !isCycle {
			fmt.Printf("%v - %v\n", fast.data, fast.data)
			isCycle = true
			cycleStart = slow
		}
	}

	return false
}
