package _69_pancake_sorting

import (
	"reflect"
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1",
			args{[]int{3, 2, 4, 1}},
			[]int{3, 4, 2, 3, 1, 2},
		}, {
			"test case 2",
			args{[]int{1, 2, 3}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pancakeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pancakeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
