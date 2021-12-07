package _034_coloring_a_border

/*
1034. 边界着色

给你一个大小为 m x n 的整数矩阵 grid ，表示一个网格。另给你三个整数 row、col 和 color 。网格中的每个值表示该位置处的网格块的颜色。

两个网格块属于同一 连通分量 需满足下述全部条件：

两个网格块颜色相同
在上、下、左、右任意一个方向上相邻
连通分量的边界 是指连通分量中满足下述条件之一的所有网格块：

在上、下、左、右四个方向上与不属于同一连通分量的网格块相邻
在网格的边界上（第一行/列或最后一行/列）
请你使用指定颜色 color 为所有包含网格块 grid[row][col] 的 连通分量的边界 进行着色，并返回最终的网格 grid 。
*/

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = grid[i][j]
		}
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[row][col] = true
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		visited[x][y] = true
		if x == 0 || x == m-1 || y == 0 || y == n-1 {
			ans[x][y] = color
		}
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				if !visited[nx][ny] && grid[nx][ny] == grid[x][y] {
					dfs(nx, ny)
				}
				if grid[nx][ny] != grid[x][y] {
					ans[x][y] = color
				}
			}
		}
	}
	dfs(row, col)
	return ans
}
