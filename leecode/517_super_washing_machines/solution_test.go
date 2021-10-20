package _17_super_washing_machines

import "testing"

func Test_findMinMoves(t *testing.T) {
	type args struct {
		machines []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{1, 0, 5},
			},
			3,
		}, {
			"test case 2",
			args{
				[]int{0, 3, 0},
			},
			2,
		}, {
			"test case 3",
			args{
				[]int{0, 2, 0},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinMoves(tt.args.machines); got != tt.want {
				t.Errorf("findMinMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
