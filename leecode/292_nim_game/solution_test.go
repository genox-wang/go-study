package nimgame

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{
				4,
			},
			false,
		},
		{
			"test case 2",
			args{
				1,
			},
			true,
		},
		{
			"test case 3",
			args{
				2,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
