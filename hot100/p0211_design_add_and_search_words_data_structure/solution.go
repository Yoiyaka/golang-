package p0211_design_add_and_search_words_data_structure

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (w *WordDictionary) AddWord(word string) {
	node := w
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionary{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (w *WordDictionary) Search(word string) bool {
	return w.searchHelper(word, 0)
}

func (w *WordDictionary) searchHelper(word string, i int) bool {
	if i == len(word) {
		return w.isEnd
	}
	c := rune(word[i])
	if c == '.' {
		for _, child := range w.children {
			if child != nil && child.searchHelper(word, i+1) {
				return true
			}
		}
		return false
	}
	idx := c - 'a'
	if w.children[idx] == nil {
		return false
	}
	return w.children[idx].searchHelper(word, i+1)
}
