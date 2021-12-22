package _86_repeated_string_march

import "testing"

func Test_repeatedStringMatch(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				"abcd",
				"cdabcdab",
			},
			3,
		}, {
			"test case 2",
			args{
				"a",
				"aa",
			},
			2,
		}, {
			"test case 1",
			args{
				"aa",
				"a",
			},
			1,
		}, {
			"test case 2",
			args{
				"aaaaaaaaaaaaaab",
				"ba",
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedStringMatch(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("repeatedStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
