package _218_longest_arithmetic_subsequence_of_given_difference

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		arr  []int
		diff int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{1, 2, 3, 4},
				1,
			},
			4,
		}, {
			"test case 2",
			args{
				[]int{1, 3, 5, 7},
				1,
			},
			1,
		}, {
			"test case 3",
			args{
				[]int{1, 5, 7, 8, 5, 3, 4, 2, 1},
				-2,
			},
			4,
		}, {
			"test case 4",
			args{
				[]int{3, 4, -3, -2, -4},
				-5,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.arr, tt.args.diff); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
