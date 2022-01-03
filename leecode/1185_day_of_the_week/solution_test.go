package _185_day_of_the_week

import "testing"

func Test_dayOfTheWeek(t *testing.T) {
	type args struct {
		day   int
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{31, 8, 2019},
			"Saturday",
		}, {
			"test case 2",
			args{18, 7, 1999},
			"Sunday",
		}, {
			"test case 3",
			args{15, 8, 1993},
			"Sunday",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfTheWeek(tt.args.day, tt.args.month, tt.args.year); got != tt.want {
				t.Errorf("dayOfTheWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
