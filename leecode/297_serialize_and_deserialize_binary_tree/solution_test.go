package _97_serialize_and_deserialize_binary_tree

import "fmt"

func ExampleConstructor() {
	ser := Constructor()
	deser := Constructor()
	ts := []string{
		"1,2,3,null,null,4,5",
		"",
		"1",
		"1,2",
	}
	for _, t := range ts {
		ans := deser.Deserialize(t)
		data := ser.Serialize(ans)
		fmt.Println(data)
	}

	// Output:
	// 1,2,3,null,null,4,5
	//
	// 1
	// 1,2
}
