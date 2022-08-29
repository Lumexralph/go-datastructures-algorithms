package isunique

import "testing"

func Test_isUnique(t *testing.T) {
	testCases := []struct {
		name, word string
		want       bool
	}{
		{
			name: "empty string",
			word: "",
			want: false,
		},
		{
			name: "unique characters",
			word: "game",
			want: true,
		},
		{
			name: "longer unique characters",
			word: "abcdefghijklm",
			want: true,
		},
		{
			name: "non-unique characters",
			word: "gayme",
			want: true,
		},
	}

	for _, tc := range testCases {
		got := isUnique(tc.word)
		if got != tc.want {
			t.Errorf("isUnique(%q): got %t; want %t", tc.word, got, tc.want)
		}
	}
}

func Test_isUniqueV2(t *testing.T) {
	testCases := []struct {
		name, word string
		want       bool
	}{
		{
			name: "empty string",
			word: "",
			want: false,
		},
		{
			name: "unique characters",
			word: "game",
			want: true,
		},
		{
			name: "longer unique characters",
			word: "abcdefghijklm",
			want: true,
		},
		{
			name: "non-unique characters",
			word: "gayme",
			want: true,
		},
	}

	for _, tc := range testCases {
		got := isUniqueV2(tc.word)
		if got != tc.want {
			t.Errorf("isUnique(%q): got %t; want %t", tc.word, got, tc.want)
		}
	}
}
