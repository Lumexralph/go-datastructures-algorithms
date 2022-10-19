package v2

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	argList := New(1)
	argList.insert(1)
	argList.insert(2)
	argList.insert(2)
	argList.insert(3)
	argList.insert(3)
	argList.insert(4)

	want := New(1)
	want.insert(2)
	want.insert(3)
	want.insert(4)

	type args struct {
		list *LinkedList
	}
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{
			name: "linked list with duplicates",
			args: args{
				list: argList,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kthToLast(t *testing.T) {
	argList := New(1)
	argList.insert(3)
	argList.insert(2)
	argList.insert(5)
	argList.insert(6)
	argList.insert(8)
	argList.insert(7)

	type args struct {
		llist    *LinkedList
		position int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4th to the last element",
			args: args{
				llist:    argList,
				position: 4,
			},
			want: 5,
		},
		{
			name: "second to the last element",
			args: args{
				llist:    argList,
				position: 2,
			},
			want: 8,
		},
		{
			name: "the last element",
			args: args{
				llist:    argList,
				position: 1,
			},
			want: 7,
		},
		{
			name: "the last element at position = 0",
			args: args{
				llist:    argList,
				position: 0,
			},
			want: 7,
		},
		{
			name: "second to the last element in a nil list",
			args: args{
				llist:    nil,
				position: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast(tt.args.llist, tt.args.position); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_kthToLastRecursion(t *testing.T) {
//	argList := New(1)
//	argList.insert(3)
//	argList.insert(2)
//	argList.insert(5)
//	argList.insert(6)
//	argList.insert(8)
//	argList.insert(7)
//
//	type args struct {
//		head     *Node
//		position int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		{
//			name: "4th to the last element",
//			args: args{
//				head:     argList.head,
//				position: 4,
//			},
//			want: 5,
//		},
//		{
//			name: "second to the last element",
//			args: args{
//				head:     argList.head,
//				position: 2,
//			},
//			want: 8,
//		},
//		{
//			name: "the last element",
//			args: args{
//				head:     argList.head,
//				position: 1,
//			},
//			want: 7,
//		},
//		{
//			name: "the last element at position = 0",
//			args: args{
//				head:     argList.head,
//				position: 0,
//			},
//			want: 7,
//		},
//		{
//			name: "second to the last element in a nil list",
//			args: args{
//				head:     nil,
//				position: 2,
//			},
//			want: 0,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := kthToLastRecursion(tt.args.head, tt.args.position); got != tt.want {
//				t.Errorf("kthToLastRecursion() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_partitionList(t *testing.T) {
	argList := New(3)
	argList.insert(5)
	argList.insert(8)
	argList.insert(5)
	argList.insert(10)
	argList.insert(2)
	argList.insert(1)

	want := New(3)
	want.insert(2)
	want.insert(1)
	want.insert(5)
	want.insert(8)
	want.insert(5)
	want.insert(10)

	type args struct {
		node *Node
		val  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "partition at value in the list",
			args: args{
				node: argList.head,
				val:  5,
			},
			want: want.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionList(tt.args.node, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				node := got
				for node.next != nil {
					t.Log("node: ", node.data)
					node = node.next
				}
				t.Errorf("partitionList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listSum(t *testing.T) {
	right := New(1)
	right.insert(5)
	right.insert(8)

	left := New(3)
	left.insert(2)
	left.insert(1)

	type args struct {
		left  *LinkedList
		right *LinkedList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum of same list length",
			args: args{
				left:  left,
				right: right,
			},
			want: 479,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listSum(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("listSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listSumReverseDigit(t *testing.T) {
	right := New(1)
	right.insert(5)
	right.insert(8)

	left := New(3)
	left.insert(2)
	left.insert(1)

	l := New(3)
	l.insert(2)
	l.insert(1)
	l.insert(8)

	type args struct {
		left  *LinkedList
		right *LinkedList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum of same length",
			args: args{
				right: right,
				left:  left,
			},
			want: 974,
		},
		{
			name: "sum of different digit length",
			args: args{
				right: l,
				left:  left,
			},
			want: 8246,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listSumReverseDigit(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("listSumReverseDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listSumDigit(t *testing.T) {
	right := New(1)
	right.insert(5)
	right.insert(8)

	left := New(3)
	left.insert(2)
	left.insert(1)

	l := New(3)
	l.insert(2)
	l.insert(1)
	l.insert(8)

	type args struct {
		left  *LinkedList
		right *LinkedList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum of same length",
			args: args{
				right: right,
				left:  left,
			},
			want: 479,
		},
		{
			name: "sum of different digit length",
			args: args{
				right: l,
				left:  left,
			},
			want: 3539,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listSumDigit(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("listSumDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLoop(t *testing.T) {
	node := &Node{
		data: 1,
		next: nil,
	}
	cycle := &Node{
		data: 4,
		next: nil,
	}
	node.next = &Node{
		data: 2,
		next: nil,
	}
	node.next.next = &Node{
		data: 3,
		next: nil,
	}
	node.next.next.next = cycle
	node.next.next.next.next = &Node{
		data: 5,
		next: nil,
	}
	node.next.next.next.next.next = &Node{
		data: 6,
		next: nil,
	}
	node.next.next.next.next.next.next = cycle

	llist := &LinkedList{
		head: node,
	}

	argList := New(3)
	argList.insert(5)
	argList.insert(8)
	argList.insert(5)
	argList.insert(10)
	argList.insert(2)
	argList.insert(1)

	type args struct {
		llist *LinkedList
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "detect cycle in the linkedlist",
			args: args{
				llist: llist,
			},
			want: true,
		},
		{
			name: "no cycle in the linkedlist",
			args: args{
				llist: argList,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLoop(tt.args.llist); got != tt.want {
				t.Errorf("isLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
