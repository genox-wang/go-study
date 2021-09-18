package splitarrayintoconsecutivesubsequences

/**
给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个长度至少为 3 的子序列，其中每个子序列都由连续整数组成。

如果可以完成上述分割，则返回 true ；否则，返回 false 。
**/

func isPossible(nums []int) bool {
	freq := map[int]int{}
	need := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	for _, num := range nums {
		if freq[num] == 0 {
			continue
		}
		if need[num] > 0 {
			freq[num]--
			need[num]--
			need[num+1]++
		} else if freq[num] > 0 && freq[num+1] > 0 && freq[num+2] > 0 {
			freq[num]--
			freq[num+1]--
			freq[num+2]--
			need[num+3]++
		} else {
			return false
		}
	}
	return true
}
