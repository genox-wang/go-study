package _97_find_the_town_judge

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{2, [][]int{{1, 2}}},
			2,
		}, {
			"test case 2",
			args{3, [][]int{{1, 3}, {2, 3}}},
			3,
		}, {
			"test case 3",
			args{3, [][]int{{1, 3}, {2, 3}, {3, 1}}},
			-1,
		}, {
			"test case 4",
			args{3, [][]int{{1, 2}, {2, 3}}},
			-1,
		}, {
			"test case 5",
			args{4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
