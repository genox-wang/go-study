package _60_single_number_2

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
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
				[]int{1, 2, 1, 3, 2, 5},
			},
			[]int{3, 5},
		}, {
			"test case 2",
			args{
				[]int{-1, 0},
			},
			[]int{-1, 0},
		}, {
			"test case 3",
			args{
				[]int{0, 1},
			},
			[]int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
