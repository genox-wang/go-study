package _18_maximum_product_of_word_lengths

/**
318. 最大单词长度乘积

给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

*/

func maxProduct(words []string) (ans int) {
	n := len(words)
	hash := make([]int, n)
	for i, word := range words {
		for _, c := range word {
			hash[i] |= 1 << (c - 'a')
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if hash[i]&hash[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
