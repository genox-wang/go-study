package _89_maximum_sum_of_3_non_overlapping_subarrays

import (
	"reflect"
	"testing"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1",
			args{
				[]int{1, 2, 1, 2, 6, 7, 5, 1},
				2,
			},
			[]int{0, 3, 5},
		}, {
			"test case 2",
			args{
				[]int{1, 2, 1, 2, 1, 2, 1, 2, 1},
				2,
			},
			[]int{0, 2, 4},
		}, {
			"test case 2",
			args{
				[]int{4, 5, 10, 6, 11, 17, 4, 11, 1, 3},
				1,
			},
			[]int{4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumOfThreeSubarrays(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumOfThreeSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
