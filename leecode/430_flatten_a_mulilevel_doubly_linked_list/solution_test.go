package _30_flatten_a_mulilevel_doubly_linked_list

import (
	"reflect"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"test case 1",
			args{arr2Node([]int{1, 2, 3, 4, 5, 6, '#', '#', '#', 7, 8, 9, 10, '#', '#', 11, 12})},
			arr2Node([]int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6}),
		}, {
			"test case 2",
			args{arr2Node([]int{1, 2, '#', 3})},
			arr2Node([]int{1, 3, 2}),
		}, {
			"test case 3",
			args{arr2Node([]int{})},
			arr2Node([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.root); !reflect.DeepEqual(node2Arr(got), node2Arr(tt.want)) {
				t.Errorf("flatten() = %v, want %v", node2Arr(got), node2Arr(tt.want))
			}
		})
	}
}
