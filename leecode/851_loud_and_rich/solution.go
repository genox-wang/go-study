package _51_loud_and_rich

/*
851. 喧闹和富有

有一组 n 个人作为实验对象，从 0 到 n - 1 编号，其中每个人都有不同数目的钱，以及不同程度的安静值（quietness）。为了方便起见，我们将编号为 x 的人简称为 "person x "。

给你一个数组 richer ，其中 richer[i] = [ai, bi] 表示 person ai 比 person bi 更有钱。另给你一个整数数组 quiet ，其中 quiet[i] 是 person i 的安静值。richer 中所给出的数据 逻辑自恰（也就是说，在 person x 比 person y 更有钱的同时，不会出现 person y 比 person x 更有钱的情况 ）。

现在，返回一个整数数组 answer 作为答案，其中 answer[x] = y 的前提是，在所有拥有的钱肯定不少于 person x 的人中，person y 是最安静的人（也就是安静值 quiet[y] 最小的人）。
*/

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	m := make([][]int, n)
	for i := range m {
		m[i] = []int{}
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	for _, r := range richer {
		m[r[0]] = append(m[r[0]], r[1])
		if quiet[ans[r[1]]] > quiet[ans[r[0]]] {
			ans[r[1]] = ans[r[0]]
			var q []int
			q = append(q, m[r[1]]...)
			for len(q) > 0 {
				count := len(q)
				for _, idx := range q {
					if quiet[ans[idx]] > quiet[ans[r[1]]] {
						ans[idx] = ans[r[1]]
					}
					q = append(q, m[idx]...)
				}
				q = q[count:]
			}
		}
	}
	return ans
}
