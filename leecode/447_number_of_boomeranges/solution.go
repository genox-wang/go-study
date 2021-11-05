package _47_number_of_boomeranges

/**
447. 回旋镖的数量

给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，
其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。

返回平面上所有回旋镖的数量。
*/

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		m := make(map[int]int, 0)
		for _, q := range points {
			k := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			m[k]++
		}
		for _, v := range m {
			ans += v * (v - 1)
		}
	}
	return ans
}
