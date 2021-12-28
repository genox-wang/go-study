package _72_concatenated_words

import "sort"

/*
472. 连接词

给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。

连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。

*/

func findAllConcatenatedWordsInADict(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	root := &Trie{}
	for _, word := range words {
		if word == "" {
			continue
		}
		if root.dfs(word) {
			ans = append(ans, word)
		} else {
			root.Insert(word)
		}
	}
	return
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (root *Trie) Insert(word string) {
	node := root
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (root *Trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		ch -= 'a'
		node = node.children[ch]
		if node == nil {
			return false
		}
		if node.isEnd && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}
