package _2_trapping_rain_water

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	lMax, rMax := height[left], height[right]
	res := 0
	for left < right {
		lMax = max(height[left], lMax)
		rMax = max(height[right], rMax)
		if lMax < rMax {
			res += lMax - height[left]
			left++
		} else {
			res += rMax - height[right]
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
