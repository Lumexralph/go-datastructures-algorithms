package binarytree_test

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	binarytree "github.com/Lumexralph/data_structures_algorithm/dataStructures/binarySearchTree"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree_Search(t *testing.T) {
	testCases := []struct {
		name  string
		items []string
		item string
		want  string
	}{
		{"word found", []string{"ball", "game", "tea", "pad"}, "pad", "pad"},
		{"word not found", []string{"ball", "game", "tea", "pad"}, "book", ""},
		{"search in empty tree", []string{}, "book", ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bst := binarytree.ProduceTreeItems(tt.items...)
			got := bst.Search(tt.item)
			assert.Equalf(t, got, tt.want, "bst.Search(%q): got=%q; want=%q", got, tt.want)
		})
	}
}

func TestBinarySearchTree_FindMinimum(t *testing.T) {
	testCases := []struct {
		name  string
		items []string
		want  string
	}{
		{"smallest word", []string{"ball", "game", "tea", "pad"}, "ball"},
		{"smallest in empty tree", []string{}, ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bst := binarytree.ProduceTreeItems(tt.items...)
			got := bst.FindMinimum()
			assert.Equalf(t, got.Item, tt.want, "bst.FindMinimum(): got=%q; want=%q", got.Item, tt.want)
		})

	}
}

func TestBinarySearchTree_FindMaximum(t *testing.T) {
	testCases := []struct {
		name  string
		items []string
		want  string
	}{
		{"largest word", []string{"ball", "game", "tea", "pad"}, "tea"},
		{"largest in empty tree", []string{}, ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bst := binarytree.ProduceTreeItemsRecursive(tt.items...)
			got := bst.FindMaximum()
			assert.Equalf(t, got.Item, tt.want, "bst.FindMaximum(): got=%q; want=%q", got.Item, tt.want)
		})

	}
}

func TestBinarySearchTree_InOrderTraversal(t *testing.T) {
	testCases := []struct {
		name  string
		items []string
		want  []string
	}{
		{"all entries", []string{"game", "tea", "pad", "ball"}, []string{"ball", "game", "pad", "tea"}},
		{"all entries in empty tree", []string{}, []string{""}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bst := binarytree.ProduceTreeItemsRecursive(tt.items...)
			got := new([]string)
			bst.InOrderTraversal(bst.Root, binarytree.ProcessItem(got))
			diff := cmp.Diff(tt.want, *got)
			assert.Equalf(t, "", diff, "\"bst.InOrderTraversal() mismatch (-want +got):\n%s", "", diff)
		})
	}
}

func TestBinarySearchTree_PreOrderTraversal(t *testing.T) {
	testCases := []struct {
		name  string
		items []string
		want  []string
	}{
		{"all entries", []string{"lame", "tad", "pea", "came"}, []string{"lame", "came", "tad", "pea"}},
		{"all entries in empty tree", []string{}, []string{""}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bst := binarytree.ProduceTreeItemsRecursive(tt.items...)
			got := new([]string)
			bst.PreOrderTraversal(bst.Root, binarytree.ProcessItem(got))
			diff := cmp.Diff(tt.want, *got)
			assert.Equalf(t, "", diff, "bst.PreOrderTraversal() mismatch (-want +got):\n%s", "", diff)
		})

	}
}

func TestBinarySearchTree_PostOrderTraversal(t *testing.T) {
	testCases := []struct {
		name  string
		items []string
		want  []string
	}{
		{"all entries", []string{"fame", "pad", "tea", "game"}, []string{"game", "tea", "pad", "fame"}},
		{"all entries in empty tree", []string{}, []string{""}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bst := binarytree.ProduceTreeItemsRecursive(tt.items...)
			got := new([]string)
			bst.PostOrderTraversal(bst.Root, binarytree.ProcessItem(got))
			diff := cmp.Diff(tt.want, *got)
			assert.Equalf(t, "", diff, "bst.PostOrderTraversal() mismatch (-want +got):\n%s", "", diff)
		})

	}
}

func generateItems(repeat int) []string {
	items := []string{"fame", "pad", "tea", "game", "deep", "class", "gopher"}
	for i := 0; i < repeat; i++ {
		items = append(items, items...)
	}
	return items
}

func benchmarkProduceItemsRecursive(items []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		binarytree.ProduceTreeItemsRecursive(items...)
	}
}
func BenchmarkProduce7TreeItemsRecursive(b *testing.B) {
	b.StopTimer()
	items := generateItems(0)
	b.StartTimer()
	benchmarkProduceItemsRecursive(items, b)
}

func BenchmarkProduce14TreeItemsRecursive(b *testing.B) {
	b.StopTimer()
	items := generateItems(1)
	b.StartTimer()
	benchmarkProduceItemsRecursive(items, b)
}

func benchmarkProduceItems(items []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		binarytree.ProduceTreeItems(items...)
	}
}
func BenchmarkProduce7TreeItems(b *testing.B) {
	b.StopTimer()
	items := generateItems(0)
	b.StartTimer()
	benchmarkProduceItems(items, b)
}

func BenchmarkProduce14TreeItems(b *testing.B) {
	b.StopTimer()
	items := generateItems(1)
	b.StartTimer()
	benchmarkProduceItems(items, b)
}




