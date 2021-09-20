package _73_number_of_longest_increasing_subsequence

import "testing"

func Test_findNumberOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{1, 3, 5, 4, 7},
			},
			2,
		}, {
			"test case 2",
			args{
				[]int{2, 2, 2, 2, 2},
			},
			5,
		}, {
			"test case 3",
			args{
				[]int{3, 1, 2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("findNumberOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
