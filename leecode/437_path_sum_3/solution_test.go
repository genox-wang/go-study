package _37_path_sum_3

import (
	ts "go-study/leecode/297_serialize_and_deserialize_binary_tree"
	"testing"
)

func Test_pathSum(t *testing.T) {
	ser := ts.Constructor()
	type args struct {
		root      string
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{"10,5,-3,3,2,null,11,3,-2,null,1", 8},
			3,
		}, {
			"test case 2",
			args{"5,4,8,11,null,13,4,7,2,null,null,5,1", 22},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(ser.Deserialize(tt.args.root), tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
