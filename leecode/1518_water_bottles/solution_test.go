package _518_water_bottles

import "testing"

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				9, 3,
			},
			13,
		}, {
			"test case 2",
			args{
				15, 4,
			},
			19,
		}, {
			"test case 3",
			args{
				5, 5,
			},
			6,
		}, {
			"test case 4",
			args{
				2, 3,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
