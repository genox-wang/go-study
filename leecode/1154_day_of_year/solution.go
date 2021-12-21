package _154_day_of_year

import (
	"strconv"
	"strings"
)

/*
1154. 一年中的第几天

给你一个字符串date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。

通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。

*/

func dayOfYear(date string) (ans int) {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	splits := strings.Split(date, "-")
	y, _ := strconv.Atoi(splits[0])
	m, _ := strconv.Atoi(splits[1])
	d, _ := strconv.Atoi(splits[2])
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		days[1]++
	}
	for i := 0; i < m-1; i++ {
		ans += days[i]
	}
	ans += d
	return
}
