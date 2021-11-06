package _68_missiing_number

/**
268. 丢失的数字

给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
*/

func missingNumber(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		ans ^= nums[i]
		ans ^= i
	}
	ans ^= n
	return
}
