package _154_day_of_year

import "testing"

func Test_dayOfYear(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				"2019-01-09",
			},
			9,
		}, {
			"test case 2",
			args{
				"2019-02-10",
			},
			41,
		}, {
			"test case 3",
			args{
				"2003-03-01",
			},
			60,
		}, {
			"test case 4",
			args{
				"2004-03-01",
			},
			61,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfYear(tt.args.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
