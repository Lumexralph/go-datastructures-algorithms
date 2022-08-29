package v2

import (
	"reflect"
	"testing"
)

func Test_maxAverageSubArr(t *testing.T) {
	type args struct {
		arr    []int
		length int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "empty array",
			args: args{
				arr:    []int{},
				length: 3,
			},
			want: 0,
		},
		{
			name: "sub-array of length 1",
			args: args{
				arr:    []int{3},
				length: 3,
			},
			want: 0,
		},
		{
			name: "sub-array of length 1 with arr of length 2",
			args: args{
				arr:    []int{1, 12, -5, -6, 50, 3},
				length: 4,
			},
			want: 12.75,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAverageSubArr(tt.args.arr, tt.args.length); got != tt.want {
				t.Errorf("maxAverageSubArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubStr(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty string",
			args: args{
				word: "",
			},
			want: 0,
		},
		{
			name: "repeated string",
			args: args{
				word: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "repeated same string",
			args: args{
				word: "bbbbbb",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubStr(tt.args.word); got != tt.want {
				t.Errorf("longestSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateImage(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "rotate 4 x 4 px image",
			args: args{
				img: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: [][]int{
				{13, 9, 5, 1},
				{14, 6, 7, 2},
				{15, 10, 11, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateImage(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zeroMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "an empty matrix",
			args: args{},
			want: [][]int{},
		},
		{
			name: "empty inner array of  matrix ",
			args: args{
				matrix: [][]int{
					{},
					{},
				},
			},
			want: [][]int{},
		},
		{
			name: "zero the row and column",
			args: args{
				matrix: [][]int{
					{1, 0, 3, 4},
					{5, 6, 7, 0},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{9, 0, 11, 0},
				{13, 0, 15, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
