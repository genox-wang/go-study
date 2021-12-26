package _078_occurences_after_bigram

import "strings"

/*
1078. Bigram 分词

给出第一个词 first 和第二个词 second，考虑在某些文本 text 中可能以 "first second third" 形式出现的情况，其中 second 紧随 first 出现，third 紧随 second 出现。

对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。
*/

func findOcurrences(text string, first string, second string) (ans []string) {
	words := strings.Split(text, " ")
	for i := range words[:len(words)-2] {
		if words[i] == first && words[i+1] == second {
			ans = append(ans, words[i+2])
		}
	}
	return
}
