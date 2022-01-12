package _34_increasing_triplet_subsequence

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{[]int{1, 2, 3, 4, 5}},
			true,
		}, {
			"test case 2",
			args{[]int{5, 4, 3, 2, 1}},
			false,
		}, {
			"test case 3",
			args{[]int{20, 100, 10, 12, 5, 13}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
