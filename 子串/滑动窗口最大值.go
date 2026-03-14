func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	deque := []int{} // 存储下标
	res := []int{}
	for i := 0; i < len(nums); i++ {
		// 出队列：移除滑出窗口的下标
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// 保持单调递减队列，移除比当前元素小的元素
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		// 入队
		deque = append(deque, i)
		// 记录结果，窗口形成后开始记录
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}