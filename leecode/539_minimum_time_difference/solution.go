package _39_minimum_time_difference

import "sort"

/*
539. 最小时间差

给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
*/

func findMinDifference(timePoints []string) (ans int) {
	sort.Slice(timePoints, func(i, j int) bool {
		return timePoints[i] < timePoints[j]
	})

	var diff = func(t1, t2 string) (ans int) {
		h := (int(t2[0])-int(t1[0]))*10 + (int(t2[1]) - int(t1[1]))
		m := (int(t2[3])-int(t1[3]))*10 + (int(t2[4]) - int(t1[4]))
		ans = h*60 + m
		if t1 > t2 {
			ans += 60 * 24
		}
		return
	}
	ans = diff(timePoints[len(timePoints)-1], timePoints[0])
	for i := 0; i < len(timePoints)-1; i++ {
		d := diff(timePoints[i], timePoints[i+1])
		if ans > d {
			ans = d
		}
	}
	return
}
