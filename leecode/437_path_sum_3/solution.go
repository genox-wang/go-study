package _37_path_sum_3

import (
	ts "go-study/leecode/297_serialize_and_deserialize_binary_tree"
)

/**
437. 路径总和 III

给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/

func pathSum(root *ts.TreeNode, targetSum int) int {
	pre := 0
	m := make(map[int]int, 0)
	m[0] = 1
	res := 0
	var dfs func(*ts.TreeNode)
	dfs = func(node *ts.TreeNode) {
		if node == nil {
			return
		}
		pre += node.Val
		if _, ok := m[pre-targetSum]; ok {
			res += m[pre-targetSum]
		}
		m[pre]++
		dfs(node.Left)
		dfs(node.Right)
		m[pre]--
		if m[pre] == 0 {
			delete(m, pre)
		}
		pre -= node.Val
	}
	dfs(root)
	return res
}
