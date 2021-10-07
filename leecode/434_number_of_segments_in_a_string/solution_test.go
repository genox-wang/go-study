package _34_number_of_segments_in_a_string

import "testing"

func Test_countSegments(t *testing.T) {
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
				"Hello, my name is John",
			},
			5,
		}, {
			"test case 2",
			args{
				"",
			},
			0,
		}, {
			"test case 3",
			args{
				" a",
			},
			1,
		}, {
			"test case 4",
			args{
				"a  ",
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
