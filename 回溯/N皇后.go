func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // 斜线1: 行+列
	diag2 := make([]bool, 2*n-1) // 斜线2: 行-列+n-1

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := 0; i < n; i++ {
				tmp[i] = string(board[i])
			}
			res = append(res, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row+col] || diag2[row-col+n-1] {
				continue
			}
			board[row][col] = 'Q'
			cols[col], diag1[row+col], diag2[row-col+n-1] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col], diag1[row+col], diag2[row-col+n-1] = false, false, false
		}
	}
	backtrack(0)
	return res
}