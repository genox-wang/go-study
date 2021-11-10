package _95_teemo_attacking

import "testing"

func Test_findPoisonedDuration(t *testing.T) {
	type args struct {
		timeSeries []int
		duration   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{1, 4},
				2,
			},
			4,
		}, {
			"test case 2",
			args{
				[]int{1, 2},
				2,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoisonedDuration(tt.args.timeSeries, tt.args.duration); got != tt.want {
				t.Errorf("findPoisonedDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
