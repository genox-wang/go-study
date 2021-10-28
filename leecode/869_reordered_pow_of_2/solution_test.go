package _69_reordered_pow_of_2

import "testing"

func Test_reorderedPowerOf2(t *testing.T) {
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
			args{10},
			false,
		}, {
			"test case 2",
			args{1},
			true,
		}, {
			"test case 3",
			args{16},
			true,
		}, {
			"test case 4",
			args{24},
			false,
		}, {
			"test case 5",
			args{46},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderedPowerOf2(tt.args.n); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
