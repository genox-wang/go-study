package _30_flatten_a_mulilevel_doubly_linked_list

/**
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	var ns []*Node
	node := root
	for node.Next != nil || node.Child != nil {
		if node.Child != nil {
			next := node.Next
			node.Next = node.Child
			node.Next.Prev = node
			node.Child = nil
			if next != nil {
				ns = append(ns, next)
			}
		}
		node = node.Next
	}
	for i := len(ns) - 1; i >= 0; i-- {
		node.Next = ns[i]
		ns[i].Prev = node
		node = ns[i]
		for node.Next != nil {
			node = node.Next
		}
	}
	return root
}

func arr2Node(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	var parents []*Node
	findingParent, parentIdx := false, 0
	var pre *Node
	node := &Node{Val: arr[0]}
	pre = node
	for _, a := range arr[1:] {
		if a == '#' && !findingParent {
			findingParent = true
			parentIdx = 0
			continue
		}
		if findingParent {
			if a == '#' {
				parentIdx++
				continue
			} else {
				findingParent = false
				pre = parents[parentIdx]
				parents = []*Node{}
			}
		}
		node := &Node{Val: a, Prev: pre}
		if pre != nil {
			pre.Next = node
		}
		pre = node
		parents = append(parents, node)
	}
	return node
}

func node2Arr(node *Node) []int {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}
