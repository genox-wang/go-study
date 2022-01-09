package _629_slowest_key

import "testing"

func Test_slowestKey(t *testing.T) {
	type args struct {
		releaseTimes []int
		keysPressed  string
	}
	tests := []struct {
		name    string
		args    args
		wantAns byte
	}{
		{
			"test case 1",
			args{[]int{9, 29, 49, 50}, "cbcd"},
			'c',
		}, {
			"test case 2",
			args{[]int{12, 23, 36, 46, 62}, "spuda"},
			'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := slowestKey(tt.args.releaseTimes, tt.args.keysPressed); gotAns != tt.wantAns {
				t.Errorf("slowestKey() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
