package _09_to_lower_case

import "strings"

/*
709. 转换成小写字母

给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。

*/

func toLowerCase(s string) string {
	var buf strings.Builder
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c |= 32
		}
		buf.WriteRune(c)
	}
	return buf.String()
}
