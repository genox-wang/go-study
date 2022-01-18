package _39_minimum_time_difference

import "testing"

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			"test case 1",
			args{[]string{"23:59", "00:00"}},
			1,
		}, {
			"test case 2",
			args{[]string{"00:00", "23:59", "00:00"}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMinDifference(tt.args.timePoints); gotAns != tt.wantAns {
				t.Errorf("findMinDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
