func largestRectangleArea(heights []int) int {
	// 单调递增栈：存下标
	stack := make([]int, 0, len(heights)+1)

	maxArea := 0

	// 遍历到 n 时，视作高度 0 的哨兵，触发清算
	for i := 0; i <= len(heights); i++ {
		curH := 0
		if i < len(heights) {
			curH = heights[i]
		}

		// 只要当前高度更小，就不断弹栈并计算以弹出柱高为最小高度的最大矩形
		for len(stack) > 0 && curH < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[top]

			// 弹出 top 后：
			// 左边界是新的栈顶（它的右侧+1），若栈空则左边界为 0
			leftLessIdx := -1
			if len(stack) > 0 {
				leftLessIdx = stack[len(stack)-1]
			}
			width := i - leftLessIdx - 1

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}