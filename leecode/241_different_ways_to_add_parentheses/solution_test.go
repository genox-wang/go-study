package _41_different_ways_to_add_parentheses

import (
	"reflect"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1",
			args{
				"2-1-1",
			},
			[]int{2, 0},
		}, {
			"test case 2",
			args{
				"2*3-4*5",
			},
			[]int{-34, -10, -14, -10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
