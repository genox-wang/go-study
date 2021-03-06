package _99_bulls_and_cows

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{
				"1807",
				"7810",
			},
			"1A3B",
		}, {
			"test case 2",
			args{
				"1123",
				"0111",
			},
			"1A1B",
		}, {
			"test case 3",
			args{
				"1",
				"0",
			},
			"0A0B",
		}, {
			"test case 4",
			args{
				"1",
				"1",
			},
			"1A0B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
