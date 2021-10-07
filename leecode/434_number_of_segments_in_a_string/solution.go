package _34_number_of_segments_in_a_string

/**
434. 字符串中的单词数

统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。
*/

func countSegments(s string) int {
	s = s + " "
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] != ' ' && s[i+1] == ' ' {
			ans++
		}
	}
	return ans
}
