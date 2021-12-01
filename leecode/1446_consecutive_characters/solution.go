package _446_consecutive_characters

/**
1446. 连续字符

给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。

请你返回字符串的能量。
*/

func maxPower(s string) int {
	cnt, ans := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			cnt = 1
		} else {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		}
	}
	return ans
}
