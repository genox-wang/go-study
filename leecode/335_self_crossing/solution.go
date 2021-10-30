package _35_self_crossing

/**
335. 路径交叉

给你一个整数数组 distance 。

从 X-Y 平面上的点 (0,0) 开始，先向北移动 distance[0] 米，然后向西移动 distance[1] 米，向南移动 distance[2] 米，向东移动 distance[3] 米，持续移动。也就是说，每次移动后你的方位会发生逆时针变化。

判断你所经过的路径是否相交。如果相交，返回 true ；否则，返回 false 。
*/

func isSelfCrossing(d []int) bool {
	if len(d) < 4 {
		return false
	}
	for i := 3; i < len(d); i++ {
		if d[i-1] <= d[i-3] && d[i] >= d[i-2] {
			return true
		} else if i >= 4 && d[i-1] == d[i-3] && d[i-4]+d[i] >= d[i-2] {
			return true
		} else if i >= 5 && d[i-2] > d[i-4] && d[i]+d[i-4] >= d[i-2] &&
			d[i-3] <= d[i-5]+d[i-1] && d[i-3] >= d[i-1] {
			return true
		}
	}
	return false
}
