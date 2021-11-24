package _84_shuffle_an_array

import (
	"math/rand"
)

/**
384. 打乱数组

给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。

实现 Solution class:

Solution(int[] nums) 使用整数数组 nums 初始化对象
int[] reset() 重设数组到它的初始状态并返回
int[] shuffle() 返回数组随机打乱后的结果
*/

type Solution struct {
	nums, original []int
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int(nil), nums...)}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)
	return s.nums
}

func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	for i := range s.nums {
		j := i + rand.Intn(n-i)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
	return s.nums
}
