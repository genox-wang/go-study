package _01_remove_invalid_parentheses

/**
301. 删除无效的括号

给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回。
*/

func removeInvalidParentheses(s string) (ans []string) {
	var dfs func(s string, i int, cnt int)
	m := make(map[int]map[string]struct{}, 0)
	maxLen := 0
	dfs = func(s string, i int, cnt int) {
		if cnt < 0 {
			return
		}
		if len(s) == i {
			if cnt == 0 {
				if maxLen <= len(s) {
					maxLen = len(s)
					if _, ok := m[maxLen]; !ok {
						m[maxLen] = make(map[string]struct{}, 0)
					}
					m[maxLen][s] = struct{}{}
				}
			}
			return
		}
		if s[i] == '(' {
			dfs(s, i+1, cnt+1)
			dfs(s[:i]+s[i+1:], i, cnt)
		} else if s[i] == ')' {
			dfs(s, i+1, cnt-1)
			dfs(s[:i]+s[i+1:], i, cnt)
		} else {
			dfs(s, i+1, cnt)
		}
	}
	dfs(s, 0, 0)
	for k := range m[maxLen] {
		ans = append(ans, k)
	}
	return
}
