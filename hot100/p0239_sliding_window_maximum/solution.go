package p0239_sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{} // stores indices
	result := []int{}
	for i, n := range nums {
		// remove elements outside window
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		// remove smaller elements
		for len(deque) > 0 && nums[deque[len(deque)-1]] < n {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}
