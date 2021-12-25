package _609_even_odd_tree

/*
1609. 奇偶树

如果一棵二叉树满足下述几个条件，则可以称为 奇偶树 ：

二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
给你二叉树的根节点，如果二叉树为 奇偶树 ，则返回 true ，否则返回 false 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	i := 0
	for len(q) > 0 {
		count := len(q)
		pre := -1
		for count > 0 {
			count--
			n := q[0]
			q = q[1:]
			if n.Val%2 == i%2 {
				return false
			}
			if pre != -1 {
				if i%2 == 2 && pre >= n.Val {
					return false
				}
				if i%2 == 1 && pre <= n.Val {
					return false
				}
			}
			pre = n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		i++
	}
	return true
}
