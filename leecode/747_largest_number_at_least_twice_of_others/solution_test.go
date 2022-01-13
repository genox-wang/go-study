package _47_largest_number_at_least_twice_of_others

import "testing"

func Test_dominantIndex(t *testing.T) {
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
			args{[]int{3, 6, 1, 0}},
			1,
		}, {
			"test case 2",
			args{[]int{1, 2, 3, 4}},
			-1,
		}, {
			"test case 3",
			args{[]int{0, 0, 3, 2}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
