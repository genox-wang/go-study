package _01_remove_invalid_parentheses

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test case 1",
			args{
				"()())()",
			},
			[]string{"()()()", "(())()"},
		}, {
			"test case 2",
			args{
				"(a)())()",
			},
			[]string{"(a)()()", "(a())()"},
		}, {
			"test case 3",
			args{
				"')('",
			},
			[]string{"''"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
