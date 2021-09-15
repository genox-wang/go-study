package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	a int
	b int
}

func main() {
	w := &W{}

	fmt.Println(w.a, w.b)

	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))

	*((*int)(b)) = 10

	fmt.Println(w.a, w.b)
}
