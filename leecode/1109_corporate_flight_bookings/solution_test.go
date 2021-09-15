package _109_corporate_flight_bookings

import (
	"reflect"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	type args struct {
		bookings [][]int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test case 1",
			args{
				[][]int{
					{1, 2, 10},
					{2, 3, 20},
					{2, 5, 25},
				},
				5,
			},
			[]int{10, 55, 45, 25, 25},
		}, {
			"test case 2",
			args{
				[][]int{
					{1, 2, 10},
					{2, 2, 15},
				},
				2,
			},
			[]int{10, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
