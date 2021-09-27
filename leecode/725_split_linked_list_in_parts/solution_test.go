package _25_split_linked_list_in_parts

import (
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test case 1",
			args{
				arr2ListNode([]int{1, 2, 3}),
				5,
			},
			[][]int{{1}, {2}, {3}, {}, {}},
		},
		{
			"test case 2",
			args{
				arr2ListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				3,
			},
			[][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(listNode2Arr(got), tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func arr2ListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	node := head
	for _, a := range arr[1:] {
		node.Next = &ListNode{Val: a}
		node = node.Next
	}
	return head
}

func listNode2Arr(nodes []*ListNode) [][]int {
	var res = make([][]int, 0)
	for _, node := range nodes {
		var r = make([]int, 0)
		for node != nil {
			r = append(r, node.Val)
			node = node.Next
		}
		res = append(res, r)
	}
	return res
}
