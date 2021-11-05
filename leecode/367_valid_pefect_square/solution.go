package _67_valid_pefect_square

/**
367. 有效的完全平方数

给定一个 正整数 num ，编写一个函数，如果 num 是一个完全平方数，则返回 true ，否则返回 false 。

进阶：不要 使用任何内置的库函数，如 sqrt 。
*/

func isPerfectSquare(num int) bool {
	l, r := 1, 2<<31-1
	for l <= r {
		mid := l + (r-l)/2
		square := mid * mid
		if num == square {
			return true
		}
		if square < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
