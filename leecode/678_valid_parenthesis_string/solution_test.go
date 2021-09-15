package _78_valid_parenthesis_string

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{"()"},
			true,
		}, {
			"test case 2",
			args{"(*)"},
			true,
		}, {
			"test case 3",
			args{"(*))"},
			true,
		}, {
			"test case 4",
			args{"("},
			false,
		}, {
			"test case 5",
			args{"**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))"},
			true,
		}, {
			"test case 6",
			args{"**))(("},
			false,
		}, {
			"test case 7",
			args{"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
