package _72_concatenated_words

import (
	"reflect"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
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
				[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			},
			[]string{"catsdogcats", "dogcatsdog", "ratcatdogcat"},
		}, {
			"test case 2",
			args{
				[]string{"cat", "dog", "catdog"},
			},
			[]string{"catdog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAllConcatenatedWordsInADict(tt.args.words); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAllConcatenatedWordsInADict() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
