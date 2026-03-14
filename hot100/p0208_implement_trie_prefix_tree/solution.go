package p0208_implement_trie_prefix_tree

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}
