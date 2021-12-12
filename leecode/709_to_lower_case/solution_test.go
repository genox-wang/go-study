package _09_to_lower_case

import "testing"

func Test_toLowerCase(t *testing.T) {
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
			args{
				"Hello",
			},
			"hello",
		}, {
			"test case 2",
			args{
				"here",
			},
			"here",
		}, {
			"test case 3",
			args{
				"LOVELY",
			},
			"lovely",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
