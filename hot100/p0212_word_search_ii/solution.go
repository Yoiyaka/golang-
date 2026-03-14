package p0212_word_search_ii

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}
			node = node.children[idx]
		}
		node.word = w
	}

	rows, cols := len(board), len(board[0])
	result := []string{}

	var dfs func(node *TrieNode, r, c int)
	dfs = func(node *TrieNode, r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] == '#' {
			return
		}
		ch := board[r][c]
		next := node.children[ch-'a']
		if next == nil {
			return
		}
		if next.word != "" {
			result = append(result, next.word)
			next.word = ""
		}
		board[r][c] = '#'
		dfs(next, r+1, c)
		dfs(next, r-1, c)
		dfs(next, r, c+1)
		dfs(next, r, c-1)
		board[r][c] = ch
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			dfs(root, r, c)
		}
	}
	return result
}
