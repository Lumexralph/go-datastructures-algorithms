package graph

import "testing"

func Test_depthFirstSearch(t *testing.T) {
	g := newGraph()
	n1 := &Node{
		name: "0",
	}
	g.insertNode(n1)
	g.insertAdjacent("0", "1")
	g.insertAdjacent("0", "4")
	g.insertAdjacent("0", "5")
	g.insertAdjacent("1", "4")
	g.insertAdjacent("1", "3")
	g.insertAdjacent("3", "4")
	g.insertAdjacent("3", "2")
	g.insertAdjacent("2", "1")

	type args struct {
		root *Node
		key  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "find node 2",
			args: args{
				root: n1,
				key:  "2",
			},
			want: true,
		},
		{
			name: "missing node 6",
			args: args{
				root: n1,
				key:  "6",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depthFirstSearch(tt.args.root, tt.args.key); got != tt.want {
				t.Errorf("depthFirstSearch() = %v, want %v, %v", got, tt.want, g)
			}
		})
	}
}
