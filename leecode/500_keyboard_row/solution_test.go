package _00_keyboard_row

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			"test case 1",
			args{
				[]string{"Hello", "Alaska", "Dad", "Peace"},
			},
			[]string{"Alaska", "Dad"},
		}, {
			"test case 2",
			args{
				[]string{"adsdf", "sfd"},
			},
			[]string{"adsdf", "sfd"},
		}, {
			"test case 3",
			args{
				[]string{"abdfs", "cccd", "a", "qwwewm"},
			},
			[]string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findWords(tt.args.words); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findWords() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
