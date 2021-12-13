package _07_max_increase_to_keep_city_skyline

/*
807. 保持城市天际线

在二维数组grid中，grid[i][j]代表位于某处的建筑物的高度。 我们被允许增加任何数量（不同建筑物的数量可能不同）的建筑物的高度。 高度 0 也被认为是建筑物。

最后，从新数组的所有四个方向（即顶部，底部，左侧和右侧）观看的“天际线”必须与原始数组的天际线相同。 城市的天际线是从远处观看时，由所有建筑物形成的矩形的外部轮廓。 请看下面的例子。

建筑物高度可以增加的最大总和是多少？
*/

func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	skylineH := make([]int, n)
	skylineV := make([]int, m)
	for i, g := range grid {
		for j, h := range g {
			if skylineV[i] < h {
				skylineV[i] = h
			}
			if skylineH[j] < h {
				skylineH[j] = h
			}
		}
	}
	for i, g := range grid {
		for j, h := range g {
			maxHeight := skylineV[i]
			if maxHeight > skylineH[j] {
				maxHeight = skylineH[j]
			}
			ans += maxHeight - h
		}
	}
	return
}
