package _73_integer_to_english_words

import "strings"

/**
273. 整数转换英文表示

将非负整数 num 转换为其对应的英文表示。
*/

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	wordsTo19 := []string{"", "One", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
		"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	wordsTens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy",
		"Eighty", "Ninety"}
	wordsThousands := []string{"", "Thousand", "Million", "Billion"}
	sb := strings.Builder{}
	toWords := func(num int) {
		if num >= 100 {
			sb.WriteString(wordsTo19[num/100])
			sb.WriteString(" Hundred ")
			num %= 100
		}
		if num >= 20 {
			sb.WriteString(wordsTens[num/10])
			sb.WriteByte(' ')
			num %= 10
		}
		if num > 0 {
			sb.WriteString(wordsTo19[num])
			sb.WriteByte(' ')
		}
	}

	for i, unit := 3, int(1e9); i >= 0; i-- {
		if curNum := num / unit; curNum > 0 {
			num %= unit
			toWords(curNum)
			sb.WriteString(wordsThousands[i])
			sb.WriteByte(' ')
		}
		unit /= 1000
	}
	return strings.TrimSpace(sb.String())
}
