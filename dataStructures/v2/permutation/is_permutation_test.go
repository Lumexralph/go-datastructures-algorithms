package permutation

import "testing"

func TestIsPermutation(t *testing.T) {
	testCases := []struct {
		name, firstWord, secondWord string
		want                        bool
	}{
		{
			name:       "same character and order",
			firstWord:  "abba",
			secondWord: "abba",
			want:       false,
		},
		{
			name:       "same character and different order",
			firstWord:  "abba",
			secondWord: "baba",
			want:       true,
		},
		{
			name:       "same character and different length",
			firstWord:  "abba",
			secondWord: "babab",
			want:       false,
		},
		{
			name:       "empty strings",
			firstWord:  "",
			secondWord: "",
			want:       false,
		},
	}

	for _, tc := range testCases {
		got := isPermutation(tc.firstWord, tc.secondWord)
		if got != tc.want {
			t.Errorf("isPermutation(%q, %q): got %t; want %t", tc.firstWord, tc.secondWord, got, tc.want)
		}
	}
}
