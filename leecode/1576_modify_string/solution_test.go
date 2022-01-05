package _576_modify_string

import "testing"

func Test_modifyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{"?zs"},
			"azs",
		}, {
			"test case 2",
			args{"ubv?w"},
			"ubvaw",
		}, {
			"test case 3",
			args{"j?qg??b"},
			"jaqgacb",
		}, {
			"test case 4",
			args{"??yw?ipkj?"},
			"acywaipkja",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyString(tt.args.s); got != tt.want {
				t.Errorf("modifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
