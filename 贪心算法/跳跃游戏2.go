func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	steps := 0
	end := 0
	far := 0
	for i := 0; i < n-1; i++ {
		// 如果当前位置都已经超过目前最远可达位置，直接不可达
		if i > far {
			return -1
		}
		if i+nums[i] > far {
			far = i + nums[i]
		}
		// 到达当前层边界，需要进行一次“跳跃层数”的推进
		if i == end {
			// 边界无法推进，说明卡死（无法到达更远的位置）
			if far == end {
				return -1
			}
			steps++
			end = far
			// 已经覆盖到终点，可以提前结束
			if end >= n-1 {
				return steps
			}
		}
	}
	// 理论上循环里已提前返回；这里兜底
	if far < n-1 {
		return -1
	}
	return steps
}