package _24_longest_word_in_dictionary_through_deleting

/**
524. 通过删除字母匹配到字典里最长单词

给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，
该字符串可以通过删除 s 中的某些字符得到。

如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
*/

func findLongestWord(s string, dictionary []string) string {
	res := ""
	for _, d := range dictionary {
		if (len(d) > len(res) ||
			(len(d) == len(res) && d < res)) &&
			valid(s, d) {
			res = d
		}
	}
	return res
}

func valid(s, d string) bool {
	i, j := 0, 0
	for i < len(s) {
		if s[i] == d[j] {
			i, j = i+1, j+1
		} else {
			i++
		}
		if j == len(d) {
			return true
		}
	}
	return false
}
