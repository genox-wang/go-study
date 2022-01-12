package _036_escape_a_large_maze

import (
	"fmt"
	"math"
)

/*
1036. 逃离大迷宫

在一个 106 x 106 的网格中，每个网格上方格的坐标为 (x, y) 。

现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty] 。数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。

每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表 blocked 上。同时，不允许走出网格。

只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true。否则，返回 false。
*/

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	visited := make(map[string]bool)
	n := len(blocked)
	max := n * (n - 1) / 2
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	maxLen := int(math.Pow10(6))
	isBlocked := func(x, y int) bool {
		for _, b := range blocked {
			if b[0] == x && b[1] == y {
				return true
			}
		}
		return false
	}
	for i := 0; i < 2; i++ {
		cnt := 1
		q := [][]int{source}
		if i == 1 {
			q = [][]int{target}
		}
		visited[fmt.Sprintf("%d_%d", q[0][0], q[0][1])] = true
		for len(q) > 0 && cnt <= max {
			count := len(q)
			for count > 0 {
				count--
				node := q[0]
				q = q[1:]
				for _, dir := range dirs {
					nx, ny := node[0]+dir[0], node[1]+dir[1]
					if nx == target[0] && ny == target[1] {
						return true
					}
					if nx >= 0 && nx < maxLen && ny >= 0 && ny < maxLen &&
						!isBlocked(nx, ny) &&
						!visited[fmt.Sprintf("%d_%d", nx, ny)] {
						cnt++
						visited[fmt.Sprintf("%d_%d", nx, ny)] = true
						q = append(q, []int{nx, ny})
					}
				}
			}
		}
		if cnt <= max {
			return false
		}
	}
	return true
}
