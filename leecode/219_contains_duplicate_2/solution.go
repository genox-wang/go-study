package _19_contains_duplicate_2

/*
219. 存在重复元素 II

给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if m[v] != 0 && i+1-m[v] <= k {
			return true
		}
		m[v] = i + 1
	}
	return false
}
