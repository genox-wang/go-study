package _610_maximum_number_of_visible_points

import (
	"math"
	"sort"
)

/*
1610. 可见点的最大数目

给你一个点数组 points 和一个表示角度的整数 angle ，你的位置是 location ，其中 location = [posx, posy] 且 points[i] = [xi, yi] 都表示 X-Y 平面上的整数坐标。

最开始，你面向东方进行观测。你 不能 进行移动改变位置，但可以通过 自转 调整观测角度。换句话说，posx 和 posy 不能改变。你的视野范围的角度用 angle 表示， 这决定了你观测任意方向时可以多宽。设 d 为你逆时针自转旋转的度数，那么你的视野就是角度范围 [d - angle/2, d + angle/2] 所指示的那片区域。
*/

func visiblePoints(points [][]int, angle int, location []int) int {
	var polarDegrees []float64
	sameCnt := 0
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees,
				math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}
	maxCnt := 0
	right := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		for right < 2*n && polarDegrees[right] <= deg+degree {
			right++
		}
		if right-i > maxCnt {
			maxCnt = right - i
		}
	}
	return maxCnt + sameCnt
}
