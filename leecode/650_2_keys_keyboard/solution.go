package _50_2_keys_keyboard

/**
650. 只有两个键的键盘

最初记事本上只有一个字符 'A' 。你每次可以对这个记事本进行两种操作：

Copy All（复制全部）：复制这个记事本中的所有字符（不允许仅复制部分字符）。
Paste（粘贴）：粘贴 上一次 复制的字符。
给你一个数字 n ，你需要使用最少的操作次数，在记事本上输出 恰好 n 个 'A' 。返回能够打印出 n 个 'A' 的最少操作次数。
*/

func minSteps(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			n /= i
			res += i
		}
	}
	return res
}
