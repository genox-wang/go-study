package _9_divide_two_integers

import "math"

/**
29. 两数相除

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
*/

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := div(dividend, divisor)
	if sign > 0 {
		return res
	}
	return -res
}

func div(a int, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tmp := b
	for tmp+tmp <= a {
		count = count + count
		tmp = tmp + tmp
	}
	return count + div(a-tmp, b)
}
