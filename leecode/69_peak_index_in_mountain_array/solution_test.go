package _9_peak_index_in_mountain_array

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				[]int{0, 1, 0},
			},
			1,
		}, {
			"test case 2",
			args{
				[]int{1, 3, 5, 4, 2},
			},
			2,
		}, {
			"test case 3",
			args{
				[]int{0, 10, 5, 2},
			},
			1,
		}, {
			"test case 4",
			args{
				[]int{3, 4, 5, 1},
			},
			2,
		}, {
			"test case 5",
			args{
				[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			},
			2,
		}, {
			"test case 5",
			args{
				[]int{18, 29, 38, 59, 98, 100, 99, 98, 90},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
