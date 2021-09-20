package _73_number_of_longest_increasing_subsequence

/**
给定一个未排序的整数数组，找到最长递增子序列的个数。
*/

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	cnt := make([]int, n)
	maxLen, res := 0, 0
	for i, x := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			if x > y {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			res = cnt[i]
		} else if dp[i] == maxLen {
			res += cnt[i]
		}
	}
	return res
}
