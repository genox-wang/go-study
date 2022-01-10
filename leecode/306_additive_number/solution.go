package _06_additive_number

/*
306. 累加数

累加数 是一个字符串，组成它的数字可以形成累加序列。

一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。

给你一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。

说明：累加序列里的数 不会 以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
*/

func isAdditiveNumber(num string) bool {
	// 穷举所有可能性 n 前一个数字，target 前连个数的和，i 遍历到的序号, index 生成的数字序号
	var dfs func(n, target, i, index int) bool
	dfs = func(n, target, i, index int) bool {
		if i == len(num) {
			return index >= 3
		}
		for j, val := i, 0; j < len(num) && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			if index < 2 {
				if dfs(val, n+val, j+1, index+1) {
					return true
				}
			} else if val == target {
				if dfs(val, n+val, j+1, index+1) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, 0, 0, 0)
}
