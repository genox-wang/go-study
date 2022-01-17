package _220_count_vowels_permutation

/*
1220. 统计元音字母序列的数目

给你一个整数 n，请你帮忙统计一下我们可以按下述规则形成多少个长度为 n 的字符串：

字符串中的每个字符都应当是小写元音字母（'a', 'e', 'i', 'o', 'u'）
每个元音 'a' 后面都只能跟着 'e'
每个元音 'e' 后面只能跟着 'a' 或者是 'i'
每个元音 'i' 后面 不能 再跟着另一个 'i'
每个元音 'o' 后面只能跟着 'i' 或者是 'u'
每个元音 'u' 后面只能跟着 'a'
由于答案可能会很大，所以请你返回 模 10^9 + 7 之后的结果。
*/

func countVowelPermutation(n int) (ans int) {
	const mod int = 1e9 + 7
	dp := []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = []int{
			((dp[1]+dp[2])%mod + dp[4]) % mod,
			(dp[0] + dp[2]) % mod,
			(dp[1] + dp[3]) % mod,
			dp[2], (dp[3] + dp[2]) % mod,
		}
	}
	for _, v := range dp {
		ans = (ans + v) % mod
	}
	return
}
