package _12_word_search_2

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test case 1",
			args{
				[][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				[]string{"oath", "pea", "eat", "rain"},
			},
			[]string{"oath", "eat"},
		}, {
			"test case 2",
			args{
				[][]byte{
					{'a', 'b', 'c'},
					{'a', 'e', 'd'},
					{'a', 'f', 'g'},
				},
				[]string{"eaabcdgfa", "eaafgdcba"},
			},
			[]string{"eaabcdgfa", "eaafgdcba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
