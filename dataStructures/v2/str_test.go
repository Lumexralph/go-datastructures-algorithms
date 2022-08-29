package v2

import "testing"

func Test_urlify(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{
				word: "",
			},
			want: "",
		},
		{
			name: "spaces in string",
			args: args{
				word: "Mr Lumex School",
			},
			want: "Mr%20Lumex%20School",
		},
		{
			name: "no spaces in string",
			args: args{
				word: "Mr-Lumex-School",
			},
			want: "Mr-Lumex-School",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := urlify(tt.args.word); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromePermutation(t *testing.T) {
	type args struct {
		phrase string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string with palindrome permutation",
			args: args{
				phrase: "Tact Coa",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromePermutation(tt.args.phrase); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oneAway(t *testing.T) {
	type args struct {
		smallestStr string
		largestStr  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "with string of different length and 1 edit",
			args: args{
				smallestStr: "ple",
				largestStr:  "pale",
			},
			want: true,
		},
		{
			name: "with string of different length and 1 edit",
			args: args{
				smallestStr: "pale",
				largestStr:  "pales",
			},
			want: true,
		},
		{
			name: "with string of same length and 1 edit",
			args: args{
				smallestStr: "pale",
				largestStr:  "bale",
			},
			want: true,
		},
		{
			name: "with string of same length and more than 1 edit",
			args: args{
				smallestStr: "pale",
				largestStr:  "bake",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneAway(tt.args.smallestStr, tt.args.largestStr); got != tt.want {
				t.Errorf("oneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchSubString(t *testing.T) {
	type args struct {
		sentence string
		subStr   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count substring in the sentence",
			args: args{
				sentence: "dear in the fear of a gear",
				subStr:   "ear",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchSubString(tt.args.sentence, tt.args.subStr); got != tt.want {
				t.Errorf("searchSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "compress to a smaller string",
			args: args{
				str: "aabcccccaaa",
			},
			want: "a2b1c5a3",
		},
		{
			name: "compress to a original string",
			args: args{
				str: "abcd",
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.str); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringRotationV1(t *testing.T) {
	type args struct {
		phrase string
		str    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "str is greater than the phrase",
			args: args{
				phrase: "day",
				str:    "over",
			},
			want: false,
		},
		{
			name: "empty phrase",
			args: args{
				phrase: "",
				str:    "over",
			},
			want: false,
		},
		{
			name: "str not a rotation substring of phrase",
			args: args{
				phrase: "popular",
				str:    "ulpq",
			},
			want: false,
		},
		{
			name: "match rotated substring",
			args: args{
				phrase: "popular",
				str:    "ualp",
			},
			want: true,
		},
		{
			name: "match rotated substring",
			args: args{
				phrase: "waterbottle",
				str:    "erbottlewat",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringRotationV1(tt.args.phrase, tt.args.str); got != tt.want {
				t.Errorf("stringRotationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
