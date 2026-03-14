package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)

	// 栈里存的是下标 index，保持 temperatures[stack] 单调递减
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		// 当前温度比栈顶那天高 -> 栈顶那天的答案就是 i - top
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		// 当前天入栈，等待未来更高温度来结算
		stack = append(stack, i)
	}

	return ans
}
