package linked_map

type LinkedListNode struct {
	Key  int
	Val  int
	Pre  *LinkedListNode
	Next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func NewLinkedList() *LinkedList {
	head := &LinkedListNode{}
	tail := &LinkedListNode{}
	head.Next = tail
	tail.Pre = head
	return &LinkedList{head, tail, 0}
}

func (l *LinkedList) AddFirst(node *LinkedListNode) {
	first := l.head.Next
	l.head.Next = node
	node.Pre = l.head
	node.Next = first
	first.Pre = node
	l.size++
}

func (l *LinkedList) AddLast(node *LinkedListNode) {
	last := l.tail.Pre
	l.tail.Pre = node
	node.Next = l.tail
	node.Pre = last
	last.Next = node
	l.size++
}

func (l *LinkedList) Remove(node *LinkedListNode) {
	if node == nil || node.Pre == nil || node.Next == nil {
		return
	}
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	l.size--
}

func (l *LinkedList) GetFirst() *LinkedListNode {
	if l.size == 0 {
		return nil
	}
	return l.head.Next
}

func (l *LinkedList) GetLast() *LinkedListNode {
	if l.size == 0 {
		return nil
	}
	return l.tail.Pre
}

func (l *LinkedList) RemoveFirst() *LinkedListNode {
	first := l.GetFirst()
	if first == nil {
		return nil
	}
	l.head.Next = first.Next
	first.Next.Pre = l.head
	l.size--
	return first
}

func (l *LinkedList) RemoveLast() *LinkedListNode {
	last := l.GetLast()
	if last == nil {
		return nil
	}
	l.tail.Pre = last.Pre
	last.Pre.Next = l.tail
	l.size--
	return last
}

func (l *LinkedList) Size() int {
	return l.size
}

type LinkedMap struct {
	mMap  map[int]*LinkedListNode
	mList *LinkedList
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		make(map[int]*LinkedListNode, 0),
		NewLinkedList(),
	}
}

func (m *LinkedMap) AddLast(key int, value int) {
	node := &LinkedListNode{Key: key, Val: value}
	m.mMap[key] = node
	m.mList.AddLast(node)
}

func (m *LinkedMap) AddFirst(key int, value int) {
	node := &LinkedListNode{Key: key, Val: value}
	m.mMap[key] = node
	m.mList.AddFirst(node)
}

func (m *LinkedMap) Remove(key int) {
	if _, ok := m.mMap[key]; !ok {
		return
	}
	node := m.mMap[key]
	delete(m.mMap, key)
	m.mList.Remove(node)
}

func (m *LinkedMap) RemoveFirst() (key, value int, success bool) {
	first := m.mList.RemoveFirst()
	if first == nil {
		return -1, -1, false
	}
	delete(m.mMap, first.Key)
	return first.Key, first.Val, true
}

func (m *LinkedMap) Size() int {
	return m.mList.size
}
