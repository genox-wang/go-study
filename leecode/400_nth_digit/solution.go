package _00_nth_digit

/**
400. 第 N 位数字

给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位数字。
*/

func findNthDigit(n int) int {
	length := 1
	for n > 9*pow(10, length-1)*length {
		n -= 9 * pow(10, length-1) * length
		length++
	}
	s := pow(10, length-1) + n/length - 1
	n -= (n / length) * length
	if n == 0 {
		return s % 10
	}
	return ((s + 1) / pow(10, length-n)) % 10
}

func pow(a, b int) int {
	res := 1
	for ; b > 0; b /= 2 {
		if b&1 > 0 {
			res = a * res
		}
		a *= a
	}
	return res
}
