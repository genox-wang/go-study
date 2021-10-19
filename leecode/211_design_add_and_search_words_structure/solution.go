package _11_design_add_and_search_words_structure

/**
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。
*/

type WordDictionary struct {
	words map[int]map[string]bool
}

func Constructor() WordDictionary {
	return WordDictionary{make(map[int]map[string]bool, 0)}
}

func (w *WordDictionary) AddWord(word string) {
	l := len(word)
	if _, has := w.words[l]; !has {
		w.words[l] = make(map[string]bool, 0)
	}
	w.words[l][word] = true
}

func (w *WordDictionary) Search(word string) bool {
	l := len(word)
	if _, has := w.words[l]; !has {
		return false
	}
	for n := range w.words[l] {
		isMatch := true
		for i := 0; i < l; i++ {
			if word[i] == '.' {
				continue
			}
			if word[i] != n[i] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return true
		}
	}
	return false
}
