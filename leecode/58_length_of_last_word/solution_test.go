package _8_length_of_last_word

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
			args{"Hello World"},
			5,
		}, {
			"test case 2",
			args{"   fly me   to   the moon  "},
			4,
		}, {
			"test case 3",
			args{"luffy is still joyboy"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
