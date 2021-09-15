package main

/**
给定一个单词数组和一个长度maxWidth，重新排版单词，使其成为每行恰好有maxWidth个字符，且左右两端对齐的文本。

你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格' '填充，使得每行恰好有 maxWidth个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于maxWidth。
输入单词数组 words至少包含一个单词。。
*/

import (
	"bytes"
	"fmt"
)

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	for len(words) > 0 {
		size := 0
		count := 0
		stack := make([]string, 0)
		for len(words) > 0 {
			w := words[0]
			if size+len(w)+count > maxWidth {
				break
			}
			stack = append(stack, w)
			words = words[1:]
			size += len(w)
			count++
		}
		spaces := make([]int, count-1)
		restSpaceSize := maxWidth - size
		sCount := count - 1
		for restSpaceSize > 0 && sCount > 0 {
			s := restSpaceSize / sCount
			for i := 0; i < sCount; i++ {
				spaces[i] += s
			}
			restSpaceSize -= sCount * s
			sCount--
		}
		var buffer bytes.Buffer
		for i := range stack {
			buffer.WriteString(stack[i])
			if len(words) == 0 && i < len(stack)-1 {
				buffer.WriteString(" ")
			} else {
				if i < len(stack)-1 {
					for c := 0; c < spaces[i]; c++ {
						buffer.WriteString(" ")
					}
				}
			}
		}

		bufSize := buffer.Len()
		fmt.Printf("bufSize %d\n", bufSize)
		for i := bufSize; i < maxWidth; i++ {
			buffer.WriteString(" ")
		}
		res = append(res, buffer.String())
	}
	return res
}
