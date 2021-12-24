package _705_maximun_number_of_eaten_apples

import "testing"

func Test_eatenApples(t *testing.T) {
	type args struct {
		apples []int
		days   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{1, 2, 3, 5, 2},
				[]int{3, 2, 1, 4, 2},
			},
			7,
		}, {
			"test case 2",
			args{
				[]int{3, 0, 0, 0, 0, 2},
				[]int{3, 0, 0, 0, 0, 2},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eatenApples(tt.args.apples, tt.args.days); got != tt.want {
				t.Errorf("eatenApples() = %v, want %v", got, tt.want)
			}
		})
	}
}
