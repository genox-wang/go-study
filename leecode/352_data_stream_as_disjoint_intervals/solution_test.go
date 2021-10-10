package _52_data_stream_as_disjoint_intervals

import "fmt"

func ExampleConstructor() {
	obj := Constructor()
	obj.AddNum(1)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(3)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(7)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(2)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(6)
	fmt.Println(obj.GetIntervals())

	// Output:
	// [[1 1]]
	// [[1 1] [3 3]]
	// [[1 1] [3 3] [7 7]]
	// [[1 3] [7 7]]
	// [[1 3] [6 7]]
}
