package _20_detect_capital

/*

我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。

*/

func detectCapitalUse(word string) bool {
	n := len(word)
	upper := 0
	for i := 0; i < n; i++ {
		if word[i] <= 'Z' {
			upper++
		}
	}
	if upper == n {
		return true
	}
	if upper == 0 {
		return true
	}
	if upper == 1 && word[0] < 'Z' {
		return true
	}
	return false
}
