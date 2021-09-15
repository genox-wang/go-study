package linked_map

import (
	"fmt"
)

func ExampleNewLinkedList() {
	list := NewLinkedList()
	list.AddLast(&LinkedListNode{Val: 0})
	list.AddFirst(&LinkedListNode{Val: 1})
	list.AddLast(&LinkedListNode{Val: 2})
	fmt.Println(list.GetLast().Val)
	fmt.Println(list.GetFirst().Val)
	fmt.Println(list.RemoveLast().Val)
	fmt.Println(list.RemoveFirst().Val)
	list.Remove(list.GetLast())
	fmt.Println(list.RemoveLast())

	// Output:
	// 2
	// 1
	// 2
	// 1
	// <nil>
}

func ExampleLinkedMap() {
	m := NewLinkedMap()
	m.AddLast(1, 10)
	m.AddFirst(2, 5)
	//fmt.Println(m.Get(1))
	//fmt.Println(m.Get(2))
	//fmt.Println(m.GetLast())
	//fmt.Println(m.GetFirst())
	//m.AddLast(3, 8)
	//fmt.Println(m.RemoveLast())
	//fmt.Println(m.RemoveFirst())
	//m.Remove(1)
	//fmt.Println(m.RemoveLast())

	// Output:
	// 10 true
	// 5 true
	// 10 true
	// 5 true
	// 8 true
	// 5 true
	// -1 false
}
