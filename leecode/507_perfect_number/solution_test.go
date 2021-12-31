package _07_perfect_number

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{28},
			true,
		}, {
			"test case 2",
			args{6},
			true,
		}, {
			"test case 3",
			args{496},
			true,
		}, {
			"test case 4",
			args{8128},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
