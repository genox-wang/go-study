package _30_kth_smallest_element_in_a_bst

import (
	a "go-study/leecode/297_serialize_and_deserialize_binary_tree"
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	deser := a.Constructor()
	type args struct {
		root *a.TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test case 1",
			args{
				deser.Deserialize("3,1,4,null,2"),
				1,
			},
			1,
		},
		{
			"test case 2",
			args{
				deser.Deserialize("5,3,6,2,4,null,null,1"),
				3,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
