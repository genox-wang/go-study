package _73_integer_to_english_words

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{
				123,
			},
			"One Hundred Twenty Three",
		}, {
			"test case 2",
			args{
				12345,
			},
			"Twelve Thousand Three Hundred Forty Five",
		}, {
			"test case 3",
			args{
				1234567891,
			},
			"One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		}, {
			"test case 4",
			args{
				50868,
			},
			"Fifty Thousand Eight Hundred Sixty Eight",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
