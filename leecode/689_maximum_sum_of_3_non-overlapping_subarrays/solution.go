package _89_maximum_sum_of_3_non_overlapping_subarrays

import "math"

/*
689. 三个无重叠子数组的最大和

给你一个整数数组 nums 和一个整数 k ，找出三个长度为 k 、互不重叠、且 3 * k 项的和最大的子数组，并返回这三个子数组。

以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。
*/

func maxSumOfThreeSubarrays(nums []int, k int) (ans []int) {
	var sum [3]int
	var maxSum [3]int
	for i := range maxSum {
		maxSum[i] = math.MinInt32
	}
	var maxSumIdx1, maxSum2Idx1, maxSum2Idx2 int
	for i := 0; i < len(nums)-k*2; i++ {
		for j := 0; j < 3; j++ {
			sum[j] += nums[i+j*k]
		}
		if i >= k-1 {
			if sum[0] > maxSum[0] {
				maxSum[0] = sum[0]
				maxSumIdx1 = i - k + 1
			}
			if maxSum[0]+sum[1] > maxSum[1] {
				maxSum[1] = maxSum[0] + sum[1]
				maxSum2Idx1, maxSum2Idx2 = maxSumIdx1, i+1
			}
			if maxSum[1]+sum[2] > maxSum[2] {
				maxSum[2] = maxSum[1] + sum[2]
				ans = []int{maxSum2Idx1, maxSum2Idx2, i + k + 1}
			}
			for j := 0; j < 3; j++ {
				sum[j] -= nums[i-k+1+j*k]
			}
		}
	}
	return
}
