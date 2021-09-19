package _50_2_keys_keyboard

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{3},
			3,
		}, {
			"test case 2",
			args{1},
			0,
		}, {
			"test case 3",
			args{88},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
