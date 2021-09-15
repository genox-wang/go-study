package _24_longest_word_in_dictionary_through_deleting

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{
				"abpcplea",
				[]string{"ale", "apple", "monkey", "plea"},
			},
			"apple",
		}, {
			"test case 2",
			args{
				"abpcplea",
				[]string{"a", "b", "c"},
			},
			"a",
		}, {
			"test case 3",
			args{
				"abce",
				[]string{"abe", "abc"},
			},
			"abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
