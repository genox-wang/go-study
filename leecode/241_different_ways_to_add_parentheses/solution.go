package _41_different_ways_to_add_parentheses

import (
	"strconv"
)

/**
给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。
你需要给出所有可能的组合的结果。有效的运算符号包含 +,-以及*。

*/

func diffWaysToCompute(expression string) []int {
	var res []int
	for i, c := range []rune(expression) {
		if c == '+' || c == '-' || c == '*' {
			var left []int
			var right []int
			left = diffWaysToCompute(expression[:i])
			right = diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					if c == '+' {
						res = append(res, l+r)
					} else if c == '-' {
						res = append(res, l-r)
					} else if c == '*' {
						res = append(res, l*r)
					}
				}
			}
		}
	}
	if len(res) == 0 {
		r, _ := strconv.Atoi(expression)
		res = append(res, r)
	}
	return res
}
