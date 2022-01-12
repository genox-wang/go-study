package _036_escape_a_large_maze

import "testing"

func Test_isEscapePossible(t *testing.T) {
	type args struct {
		blocked [][]int
		source  []int
		target  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{
				[][]int{{0, 1}, {1, 0}},
				[]int{0, 0},
				[]int{0, 2},
			},
			false,
		}, {
			"test case 2",
			args{
				[][]int{},
				[]int{0, 0},
				[]int{999999, 999999},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEscapePossible(tt.args.blocked, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("isEscapePossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
