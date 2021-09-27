package _39_decode_ways_2

/**
一条包含字母 A-Z 的消息通过以下的方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
要 解码 一条已编码的消息，所有的数字都必须分组，然后按原来的编码方案反向映射回字母（可能存在多种方式）。例如，"11106" 可以映射为：

"AAJF" 对应分组 (1 1 10 6)
"KJF" 对应分组 (11 10 6)
注意，像 (1 11 06) 这样的分组是无效的，因为 "06" 不可以映射为 'F' ，因为 "6" 与 "06" 不同。

除了 上面描述的数字字母映射方案，编码消息中可能包含 '*' 字符，可以表示从 '1' 到 '9' 的任一数字（不包括 '0'）。例如，编码字符串 "1*" 可以表示 "11"、"12"、"13"、"14"、"15"、"16"、"17"、"18" 或 "19" 中的任意一条消息。对 "1*" 进行解码，相当于解码该字符串可以表示的任何编码消息。

给你一个字符串 s ，由数字和 '*' 字符组成，返回 解码 该字符串的方法 数目 。

由于答案数目可能非常大，返回对 109 + 7 取余 的结果。
*/

func numDecodings(s string) int {
	const mod = 1e9 + 7
	a, b, c := 0, 1, 0
	for i := 0; i < len(s); i++ {
		c = (b * digital1(s, i)) % mod
		if i > 0 {
			c = (c + a*digital2(s, i)) % mod
		}
		a, b = b, c
	}
	return c
}

func digital1(s string, i int) int {
	ch := s[i]
	if ch == '*' {
		return 9
	}
	if ch == '0' {
		return 0
	}
	return 1
}

func digital2(s string, i int) int {
	ch1 := s[i-1]
	ch2 := s[i]
	if ch1 == '*' && ch2 == '*' {
		return 15
	}
	if ch1 == '*' {
		if ch2 <= '6' {
			return 2
		}
		return 1
	}
	if ch2 == '*' {
		if ch1 == '1' {
			return 9
		} else if ch1 == '2' {
			return 6
		}
		return 0
	}
	if ch1 != '0' && (ch1-'0')*10+(ch2-'0') <= 26 {
		return 1
	}
	return 0
}