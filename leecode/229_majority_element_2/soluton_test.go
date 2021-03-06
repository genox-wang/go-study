package _29_majority_element_2

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1",
			args{
				[]int{3, 2, 3},
			},
			[]int{3},
		}, {
			"test case 2",
			args{
				[]int{1},
			},
			[]int{1},
		}, {
			"test case 3",
			args{
				[]int{1, 1, 1, 3, 3, 2, 2, 2},
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
