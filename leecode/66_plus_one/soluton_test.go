package _6_plus_one

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1",
			args{[]int{1, 2, 3}},
			[]int{1, 2, 4},
		}, {
			"test case 2",
			args{[]int{4, 3, 2, 1}},
			[]int{4, 3, 2, 2},
		}, {
			"test case 3",
			args{[]int{0}},
			[]int{1},
		}, {
			"test case 4",
			args{[]int{9}},
			[]int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
