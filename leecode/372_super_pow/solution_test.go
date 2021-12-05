package _72_super_pow

import "testing"

func Test_superPow(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				2,
				[]int{3},
			},
			8,
		}, {
			"test case 2",
			args{
				2,
				[]int{1, 0},
			},
			1024,
		}, {
			"test case 3",
			args{
				1,
				[]int{4, 3, 3, 8, 5, 2},
			},
			1,
		}, {
			"test case 4",
			args{
				2147483647,
				[]int{2, 0, 0},
			},
			1198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := superPow(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
