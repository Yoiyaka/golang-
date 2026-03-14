type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{children: make(map[byte]*TrieNode)},
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.children[c] == nil {
			node.children[c] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.children[c] == nil {
			return false
		}
		node = node.children[c]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if node.children[c] == nil {
			return false
		}
		node = node.children[c]
	}
	return true
}

/*
Your Trie object will be instantiated and called as such:
obj := Constructor();
obj.Insert(word);
param_2 := obj.Search(word);
param_3 := obj.StartsWith(prefix);
*/