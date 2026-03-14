func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	queue := [][2]int{}
	fresh := 0
	// 初始化：加入所有腐烂橘子，统计新鲜橘子数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	// BFS
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					fresh--
					queue = append(queue, [2]int{nx, ny})
				}
			}
		}
		minutes++
	}
	// 若还有新鲜橘子，则无法全部腐烂
	if fresh > 0 {
		return -1
	}
	return minutes
}