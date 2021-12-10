package _23_reconstruct_original_digits_from_english

import (
	"strings"
)

/**
423. 从英文中重建数字

给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
*/

func originalDigits(s string) (ans string) {
	cnt := make(map[byte]int)
	for i := range s {
		cnt[s[i]]++
	}
	var ns [10]int
	ns[0], ns[6], ns[4], ns[8] = cnt['z'], cnt['x'], cnt['u'], cnt['g']
	ns[5], ns[7], ns[3] = cnt['f']-ns[4], cnt['s']-ns[6], cnt['h']-ns[8]
	ns[2] = cnt['t'] - ns[3] - ns[8]
	ns[1] = cnt['o'] - ns[0] - ns[2] - ns[4]
	ns[9] = cnt['i'] - ns[5] - ns[6] - ns[8]
	var buf strings.Builder
	for i, n := range ns {
		for j := 0; j < n; j++ {
			buf.WriteByte('0' + byte(i))
		}
	}
	return buf.String()
}
