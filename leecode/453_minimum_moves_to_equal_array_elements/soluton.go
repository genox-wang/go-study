package _53_minimum_moves_to_equal_array_elements

import "math"

/**
453. 最小操作次数使数组元素相等

给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
*/

func minMoves(nums []int) int {
	sum, min := 0, math.MaxInt32
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}
	return sum - len(nums)*min
}
