package _47_largest_number_at_least_twice_of_others

import "math"

/*
747. 至少是其他数字两倍的最大数

给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。

请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。
*/

func dominantIndex(nums []int) int {
	maxI, maxVal := 0, nums[0]
	otherMax := math.MinInt32
	for i, num := range nums[1:] {
		if num > maxVal {
			maxI = i + 1
			otherMax = maxVal
			maxVal = num
		} else if num > otherMax {
			otherMax = num
		}
	}
	if otherMax*2 <= maxVal {
		return maxI
	}
	return -1
}
