package _11_design_add_and_search_words_structure

import "fmt"

func ExampleConstructor() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Println(obj.Search("pad"))
	fmt.Println(obj.Search("bad"))
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("b.."))

	// Output:
	// false
	// true
	// true
	// true
}
