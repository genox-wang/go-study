package _39_decode_ways_2

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				"*",
			},
			9,
		}, {
			"test case 2",
			args{
				"1*",
			},
			18,
		}, {
			"test case 3",
			args{
				"2*",
			},
			15,
		}, {
			"test case 4",
			args{
				"3*",
			},
			9,
		}, {
			"test case 5",
			args{
				"*1",
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
