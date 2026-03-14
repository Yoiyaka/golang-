package p0051_n_queens

func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	cols := make([]bool, n)
	diag1 := make([]bool, 2*n)
	diag2 := make([]bool, 2*n)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			for i, r := range board {
				solution[i] = string(r)
			}
			result = append(result, solution)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col+n] || diag2[row+col] {
				continue
			}
			board[row][col] = 'Q'
			cols[col] = true
			diag1[row-col+n] = true
			diag2[row+col] = true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col] = false
			diag1[row-col+n] = false
			diag2[row+col] = false
		}
	}
	backtrack(0)
	return result
}
