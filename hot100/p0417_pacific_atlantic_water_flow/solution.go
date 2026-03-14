package p0417_pacific_atlantic_water_flow

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] && heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, visited)
			}
		}
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, pacific)
		dfs(r, cols-1, atlantic)
	}
	for c := 0; c < cols; c++ {
		dfs(0, c, pacific)
		dfs(rows-1, c, atlantic)
	}
	result := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}
