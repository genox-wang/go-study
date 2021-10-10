package _41_arranging_coins

import "testing"

func Test_arrangeCoins(t *testing.T) {
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
			args{
				5,
			},
			2,
		}, {
			"test case 2",
			args{
				8,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
