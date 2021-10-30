package _60_single_number_2

/**
260. 只出现一次的数字 III

给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
*/

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	a, b := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
