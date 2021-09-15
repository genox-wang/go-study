package lru

import (
	"fmt"
)

func ExampleNewLinkedList() {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2))
	fmt.Println(lruCache.Get(3))
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))
	// Output:
	// 1
	// -1
	// 3
	// -1
	// 3
	// 4
}
