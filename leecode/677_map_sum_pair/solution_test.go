package _77_map_sum_pair

import "fmt"

func ExampleConstructor() {
	obj := Constructor()
	obj.Insert("apple", 3)
	fmt.Println(obj.Sum("ap"))
	obj.Insert("app", 2)
	fmt.Println(obj.Sum("ap"))

	// Output:
	// 3
	// 5
}
