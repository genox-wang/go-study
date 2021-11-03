package _00_keyboard_row

import (
	"strings"
)

/**
500. 键盘行

给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。

美式键盘 中：

第一行由字符 "qwertyuiop" 组成。
第二行由字符 "asdfghjkl" 组成。
第三行由字符 "zxcvbnm" 组成
*/

func findWords(words []string) (ans []string) {
	const rowIdx = "12210111011122000010020202"
next:
	for _, word := range words {
		w := strings.ToLower(word)
		idx := rowIdx[w[0]-'a']
		for _, ch := range w[1:] {
			if rowIdx[ch-'a'] != idx {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return
}
