package p0207_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
	}
	// 0=unvisited, 1=visiting, 2=visited
	state := make([]int, numCourses)
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] == 1 {
			return false // cycle
		}
		if state[node] == 2 {
			return true
		}
		state[node] = 1
		for _, next := range adj[node] {
			if !dfs(next) {
				return false
			}
		}
		state[node] = 2
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
