package _07_trapping_rain_water_2

import "container/heap"

/**
407. 接雨水 II

给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
*/

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	if m < 3 && n < 3 {
		return 0
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	h := &hp{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, cell{i, j, heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}
	dirs := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if nx > 0 && nx < m-1 && ny > 0 && ny < n-1 && !visited[nx][ny] {
				if cur.h > heightMap[nx][ny] {
					ans += cur.h - heightMap[nx][ny]
				}
				visited[nx][ny] = true
				heap.Push(h, cell{nx, ny, max(heightMap[nx][ny], cur.h)})
			}
		}
	}
	return ans
}

type cell struct{ x, y, h int }

type hp []cell

func (h hp) Less(i, j int) bool { return h[i].h < h[j].h }

func (h hp) Len() int { return len(h) }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v interface{}) { *h = append(*h, v.(cell)) }

func (h *hp) Pop() interface{} { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
