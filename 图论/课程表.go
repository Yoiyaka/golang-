func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建图和入度数组
	graph := make([][]int, numCourses) // 邻接表
	indegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		cur, preq := pre[0], pre[1]
		graph[preq] = append(graph[preq], cur)
		indegree[cur]++
	}
	// 入度为0的点入队
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	finished := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		finished++
		for _, next := range graph[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return finished == numCourses
}