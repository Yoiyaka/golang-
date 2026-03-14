func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	res := 0
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			res = max(res, dp[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*堆栈
func longestValidParentheses(s string) int {
    stack := []int{-1} // 初始放一个-1，目的是方便计算长度
    res := 0
    for i, ch := range s {
        if ch == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1] // pop
            if len(stack) == 0 {
                stack = append(stack, i) // 栈空时push当前下标，代表新的起点
            } else {
                res = max(res, i-stack[len(stack)-1])
            }
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
*/