func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// 每进一次递归就保存当前 path 到结果中
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // 加入当前数字
			backtrack(i + 1)             // 递归处理后续
			path = path[:len(path)-1]    // 回溯移除
		}
	}

	backtrack(0)
	return res
}