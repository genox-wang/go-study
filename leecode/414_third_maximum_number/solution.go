package _14_third_maximum_number

import "sort"

/**
给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
*/

func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, diff := 1, 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			diff++
		}
		if diff == 3 {
			return nums[i]
		}
	}
	return nums[0]
}
