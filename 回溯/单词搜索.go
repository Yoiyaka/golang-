func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	var dfs func(x, y, idx int) bool
	dfs = func(x, y, idx int) bool {
		if idx == len(word) {
			return true
		}
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] != word[idx] {
			return false
		}
		// 标记已经访问过的格子，防止重复使用
		tmp := board[x][y]
		board[x][y] = '#'
		found := dfs(x+1, y, idx+1) || dfs(x-1, y, idx+1) || dfs(x, y+1, idx+1) || dfs(x, y-1, idx+1)
		board[x][y] = tmp // 恢复现场
		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}