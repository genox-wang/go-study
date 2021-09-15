package _78_valid_parenthesis_string

/**
给定一个只包含三种字符的字符串：（，）和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

任何左括号 (必须有相应的右括号 )。
任何右括号 )必须有相应的左括号 (。
左括号 ( 必须在对应的右括号之前 )。
*可以被视为单个右括号 )，或单个左括号 (，或一个空字符串。
一个空字符串也被视为有效字符串。
*/

func checkValidString(s string) bool {
	lo, hi := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i : i+1]
		if c == "(" {
			lo++
			hi++
		} else if c == ")" {
			lo = max(lo-1, 0)
			hi--
			if hi < 0 {
				return false
			}
		} else {
			lo = max(lo-1, 0)
			hi++
		}
	}
	return lo <= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
