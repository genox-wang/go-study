package _185_day_of_the_week

/*
1185. 一周中的第几天

给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。

输入为三个整数：day、month 和 year，分别表示日、月、年。

您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}。
*/

func dayOfTheWeek(d int, m int, y int) string {
	if m <= 2 {
		m += 12
		y--
	}
	w := (d + 1 + 2*m + 3*(m+1)/5 + y + y/4 - y/100 + y/400) % 7
	weekDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return weekDays[w]
}
