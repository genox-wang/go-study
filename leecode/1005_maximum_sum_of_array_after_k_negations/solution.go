package _005_maximum_sum_of_array_after_k_negations

import (
	"math"
	"sort"
)

/*

1005. K 次取反后最大化的数组和
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。
*/

func largestSumAfterKNegations(nums []int, k int) (ans int) {
	sort.Ints(nums)
	minNum := math.MaxInt32
	for _, num := range nums {
		if k > 0 && num < 0 {
			k--
			num = -num
		}
		ans += num
		if minNum > num {
			minNum = num
		}
	}
	if k%2 == 1 {
		ans -= 2 * minNum
	}
	return
}
