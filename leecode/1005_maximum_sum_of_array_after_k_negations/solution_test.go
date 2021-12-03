package _005_maximum_sum_of_array_after_k_negations

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{4, 2, 3},
				1,
			},
			5,
		}, {
			"test case 2",
			args{
				[]int{3, -1, 0, 2},
				3,
			},
			6,
		}, {
			"test case 3",
			args{
				[]int{2, -3, -1, 5, -4},
				2,
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}
