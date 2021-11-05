package _30_kth_smallest_element_in_a_bst

/**
230. 二叉搜索树中第K小的元素

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
*/

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
