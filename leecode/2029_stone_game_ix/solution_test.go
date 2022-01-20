package _029_stone_game_ix

import "testing"

func Test_stoneGameIX(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{[]int{2, 1}},
			true,
		}, {
			"test case 2",
			args{[]int{2}},
			false,
		}, {
			"test case 3",
			args{[]int{5, 1, 2, 4, 3}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIX(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameIX() = %v, want %v", got, tt.want)
			}
		})
	}
}
