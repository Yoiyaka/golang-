package p0994_rotting_oranges

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{}
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		minutes++
		nextQueue := [][2]int{}
		for _, pos := range queue {
			for _, d := range dirs {
				nr, nc := pos[0]+d[0], pos[1]+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					nextQueue = append(nextQueue, [2]int{nr, nc})
				}
			}
		}
		queue = nextQueue
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
