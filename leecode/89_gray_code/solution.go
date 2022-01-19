package _9_gray_code

/*
89. 格雷编码

n 位格雷码序列 是一个由 2n 个整数组成的序列，其中：
每个整数都在范围 [0, 2n - 1] 内（含 0 和 2n - 1）
第一个整数是 0
一个整数在序列中出现 不超过一次
每对 相邻 整数的二进制表示 恰好一位不同 ，且
第一个 和 最后一个 整数的二进制表示 恰好一位不同
给你一个整数 n ，返回任一有效的 n 位格雷码序列 。

*/

func grayCode(n int) (ans []int) {
	ans = make([]int, 1, 1<<n)
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= -0; j-- {
			ans = append(ans, ans[j]|1<<i)
		}
	}
	return
}