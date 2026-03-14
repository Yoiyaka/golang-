func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	var dfs func(idx int, remain int)
	dfs = func(idx int, remain int) {
		if remain == 0 {
			// 深拷贝一份path, append时要用新的slice!
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if remain < 0 {
			return
		}
		// 从idx开始可重复选取
		for i := idx; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, remain-candidates[i])
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)
	return res
}