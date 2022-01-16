package _82_linked_list_random_node

import "math/rand"

/*
382. 链表随机节点

给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。

实现 Solution 类：

Solution(ListNode head) 使用整数数组初始化对象。
int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
	for node, i := s.head, 1; node != nil; node, i = node.Next, i+1 {
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
	}
	return
}
