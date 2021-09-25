package _83_delete_operate_for_two_strings

/**
给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，
每步可以删除任意一个字符串中的一个字符。
*/

func minDistance(word1 string, word2 string) int {
	lcs := longestCommonSequence(word1, word2)
	return len(word1) + len(word2) - 2*lcs
}

func longestCommonSequence(s1 string, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	dp := make([]int, n2+1)
	for i := 1; i <= n1; i++ {
		pre := 0
		for j := 1; j <= n2; j++ {
			tmp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
			pre = tmp
		}
	}
	return dp[n2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
