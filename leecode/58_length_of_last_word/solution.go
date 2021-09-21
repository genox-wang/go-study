package _8_length_of_last_word

/**
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
*/

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else if length != 0 {
			return length
		}
	}
	return length
}
