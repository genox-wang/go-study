package _96_next_greater_element_1

/**
496. 下一个更大元素 I

给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。

请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			m[num] = -1
		} else {
			m[num] = stack[len(stack)-1]
		}
		stack = append(stack, num)
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = m[nums1[i]]
	}
	return nums1
}
