package _82_expression_add_operators

/**
282. 给表达式添加运算符

给定一个仅包含数字 0-9 的字符串 num 和一个目标值整数 target ，
在 num 的数字之间添加 二元 运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。
*/

func addOperators(num string, target int) (ans []string) {
	n := len(num)
	var dfs func([]byte, int, int, int)
	dfs = func(expr []byte, i int, res int, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)
		}
		for j, val := i, 0; j < n && (i == j || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				dfs(expr, j+1, val, val)
			} else {
				expr[signIndex] = '+'
				dfs(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				dfs(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				dfs(expr, j+1, (res-mul)+mul*val, mul*val)
			}
		}
	}
	dfs([]byte{}, 0, 0, 0)
	return
}
