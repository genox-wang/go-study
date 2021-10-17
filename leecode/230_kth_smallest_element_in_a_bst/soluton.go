package _30_kth_smallest_element_in_a_bst

import (
	t "go-study/leecode/297_serialize_and_deserialize_binary_tree"
)

func kthSmallest(root *t.TreeNode, k int) (ans int) {
	var stack []*t.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			ans = root.Val
			break
		}
		root = root.Right
	}
	return
}
