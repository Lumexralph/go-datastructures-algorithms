package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insertionSort(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"unsorted array",
			args{input: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"duplicate integers",
			args{input: []int{0, 9, 8, 7, 6, 5, 0, 4, 3, 8, 9, 2, 1}},
			[]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertionSort(tt.args.input)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("insertionSort() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
