package p0210_course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
	}
	state := make([]int, numCourses)
	order := []int{}
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if state[node] == 1 {
			return false
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
		order = append(order, node)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}
	// reverse
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}
