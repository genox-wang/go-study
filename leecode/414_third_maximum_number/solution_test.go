package _14_third_maximum_number

import "testing"

func Test_thirdMax(t *testing.T) {
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
			args{[]int{3, 2, 1}},
			1,
		}, {
			"test case 2",
			args{[]int{1, 2}},
			2,
		}, {
			"test case 3",
			args{[]int{2, 2, 3, 1}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
