package _022_covert_1d_array_into_2d_array

import (
	"reflect"
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			"test case 1",
			args{
				[]int{1, 2, 3, 4},
				2, 2,
			},
			[][]int{{1, 2}, {3, 4}},
		}, {
			"test case 2",
			args{
				[]int{1, 2, 3},
				1, 3,
			},
			[][]int{{1, 2, 3}},
		}, {
			"test case 3",
			args{
				[]int{1, 2},
				1, 1,
			},
			[][]int{},
		}, {
			"test case 4",
			args{
				[]int{3},
				1, 2,
			},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("construct2DArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
