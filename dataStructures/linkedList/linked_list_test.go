package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	l := newNameList("a", "b", "c")
	tests := []struct {
		want string
	}{
		{"b"},
		{"a"},
		{"c"},
	}

	for _, tt := range tests {
		got := l.searchList(tt.want)

		if got == nil {
			t.Fatal("searchList(): should not have a nil value")
		}
		if got.name != tt.want {
			t.Errorf("searchList(): got=%q; want=%q", got.name, tt.want)
		}
	}

	// when the element is not present in the list
	noNodeTests := []struct {
		val string
	}{
		{"d"},
		{"e"},
		{"f"},
		{"g"},
	}

	for _, tt := range noNodeTests {
		got := l.searchList(tt.val)

		if got != nil {
			t.Errorf("searchList(): got=%v; want=%v", got, tt)
		}
	}

	// when deletion happens
	for _, tt := range tests {
		l.deleteList(tt.want)
		got := l.searchList(tt.want)

		if got != nil {
			t.Errorf("deleteList(%q): got=%v; want=%v", tt.want, got, nil)
		}
	}
}
