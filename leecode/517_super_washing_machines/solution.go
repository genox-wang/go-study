package _17_super_washing_machines

/**
假设有 n 台超级洗衣机放在同一排上。开始的时候，每台洗衣机内可能有一定量的衣服，也可能是空的。

在每一步操作中，你可以选择任意 m (1 <= m <= n) 台洗衣机，与此同时将每台洗衣机的一件衣服送到相邻的一台洗衣机。

给定一个整数数组 machines 代表从左至右每台洗衣机中的衣物数量，请给出能让所有洗衣机中剩下的衣物的数量相等的 最少的操作步数 。如果不能使每台洗衣机中衣物的数量相等，则返回 -1 。
*/

func findMinMoves(machines []int) int {
	n := len(machines)
	sum := 0
	for _, machine := range machines {
		sum += machine
	}
	if sum%n != 0 {
		return -1
	}
	t := sum / n
	l, r := 0, sum
	ans := 0
	for i, m := range machines {
		r -= m
		left := max(i*t-l, 0)
		right := max((n-i-1)*t-r, 0)
		ans = max(ans, left+right)
		l += m
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
