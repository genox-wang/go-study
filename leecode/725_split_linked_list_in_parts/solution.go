package _25_split_linked_list_in_parts

/**
给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。

每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。

这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。

返回一个由上述 k 部分组成的数组。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	q, r := n/k, n%k
	parts := make([]*ListNode, k)
	for i, cur := 0, head; cur != nil && i < k; i++ {
		parts[i] = cur
		partSize := q
		for partSize < q-1 {
			partSize++
			cur = cur.Next
		}
		if i < r {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			cur = cur.Next
		}
		cur, cur.Next = cur.Next, nil
	}
	return parts
}
