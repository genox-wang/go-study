package _72_super_pow

/*
372. 超级次方

你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
*/

const mod = 1337

func superPow(a int, b []int) int {
	ans := 1
	for _, e := range b {
		ans = pow(ans, 10) * pow(a, e) % mod
	}
	return ans
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 > 0 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}
